package twitter

import (
	"bytes"
	"context"
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

func (c *Client) GetTimeline() ([]Tweet, error) {
	response, err := c.get(
		apiURL("/1.1/statuses/home_timeline.json"),
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	return c.tweetsByResponse(response)
}

func (c *Client) GetList(ctx context.Context) {
}

func (c *Client) GetListTimeline(ctx context.Context) {
}

func (c *Client) DoTweet(ctx context.Context) {
}

func (c *Client) DoReply(ctx context.Context) {
}

func (c *Client) DoRetweet(ctx context.Context) {
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

func (c *Client) tweetsByResponse(response *http.Response) ([]Tweet, error) {
	decoder := c.jsonDecoder(response)
	tweets := []Tweet{}
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
