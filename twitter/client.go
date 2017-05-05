package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mrjones/oauth"
)

var (
	twitterAPIURL      = "https://api.twitter.com"
	requestTokenPath   = "/oauth/request_token"
	authorizeTokenPath = "/oauth/authorize"
	accessTokenPath    = "/oauth/access_token"
)

// Client for twitter.
type Client struct {
	token *oauth.AccessToken
	cons  *oauth.Consumer
}

func NewClient(accessToken string, accessTokenSecret string) *Client {
	token := &oauth.AccessToken{
		Token:  accessToken,
		Secret: accessTokenSecret,
	}
	return &Client{token: token}
}

func (c *Client) GetTimeline() ([]*Tweet, error) {
	response, err := c.get(
		apiURL("/1.1/statuses/home_timeline.json"),
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	return c.tweetsByResponse(response)
}

func (c *Client) GetMentions() ([]*Tweet, error) {
	response, err := c.get(
		apiURL("/1.1/statuses/mentions_timeline.json"),
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	return c.tweetsByResponse(response)
}

func (c *Client) GetListTimeline(screenName string, slug string) ([]*Tweet, error) {
	response, err := c.get(
		apiURL("/1.1/lists/statuses.json"),
		map[string]string{
			"owner_screen_name": screenName,
			"slug":              slug,
		},
	)
	if err != nil {
		return nil, err
	}

	return c.tweetsByResponse(response)
}

func (c *Client) DoTweet(text string) error {
	_, err := c.post(
		apiURL("/1.1/statuses/update.json"),
		map[string]string{
			"status": text,
		},
	)
	return err
}

func (c *Client) DoReply(text string, tweetId int64) error {
	_, err := c.post(
		apiURL("/1.1/statuses/update.json"),
		map[string]string{
			"status":                text,
			"in_reply_to_status_id": fmt.Sprintf("%d", tweetId),
		},
	)
	return err
}

func (c *Client) DoFavorite(tweetId int64) error {
	_, err := c.post(
		apiURL("/1.1/favorites/create.json"),
		map[string]string{
			"id": fmt.Sprintf("%d", tweetId),
		},
	)
	return err
}

func (c *Client) DoRetweet(tweetId int64) error {
	_, err := c.post(
		apiURL("/1.1/statuses/retweet/%d.json", tweetId),
		map[string]string{},
	)
	return err
}

func (c *Client) get(requestURL string, params map[string]string) (*http.Response, error) {
	return c.consumer().Get(requestURL, params, c.token)
}

func (c *Client) post(requestURL string, params map[string]string) (*http.Response, error) {
	return c.consumer().Post(requestURL, params, c.token)
}

func (c *Client) consumer() *oauth.Consumer {
	if c.cons != nil {
		return c.cons
	}

	provider := oauth.ServiceProvider{
		RequestTokenUrl:   apiURL(requestTokenPath),
		AuthorizeTokenUrl: apiURL(authorizeTokenPath),
		AccessTokenUrl:    apiURL(accessTokenPath),
	}

	credential := NewCredential()

	c.cons = oauth.NewConsumer(
		credential.ConsumerKey,
		credential.ConsumerSecret,
		provider,
	)

	return c.cons
}

func (c *Client) tweetsByResponse(response *http.Response) ([]*Tweet, error) {
	decoder := c.jsonDecoder(response)
	tweets := []*Tweet{}
	decoder.Decode(&tweets)
	return tweets, nil
}

func (c *Client) jsonDecoder(response *http.Response) *json.Decoder {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return json.NewDecoder(bytes.NewReader(data))
}

func apiURL(format string, args ...interface{}) string {
	apiPath := fmt.Sprintf(format, args...)
	return twitterAPIURL + apiPath
}
