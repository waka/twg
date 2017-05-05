package main

import (
	"errors"

	"github.com/waka/twg/twitter"
)

type Consumer struct {
	config     *Config
	screenName string
	apiClient  *twitter.Client
}

func NewConsumer() *Consumer {
	return &Consumer{config: LoadConfig()}
}

func (cons *Consumer) Auth() error {
	if cons.config.IsAuthenticated() == false {
		accessToken, err := twitter.Authenticate()
		if err != nil {
			return err
		}
		screenName := accessToken.AdditionalData["screen_name"]

		cons.config.Set(accessToken.Token, accessToken.Secret, screenName)
		if cons.config.Save() == false {
			return errors.New("Failed to save config")
		}
	}

	cons.apiClient = twitter.NewClient(
		cons.config.AccessToken,
		cons.config.AccessTokenSecret,
	)
	cons.screenName = cons.config.ScreenName

	return nil
}

func (cons *Consumer) GetTimeline() ([]*twitter.Tweet, error) {
	return cons.apiClient.GetTimeline()
}

func (cons *Consumer) GetMentions() ([]*twitter.Tweet, error) {
	return cons.apiClient.GetMentions()
}

func (cons *Consumer) GetListTimeline(listName string) ([]*twitter.Tweet, error) {
	screenName := cons.screenName
	return cons.apiClient.GetListTimeline(screenName, listName)
}

func (cons *Consumer) Tweet(text string) error {
	return cons.apiClient.DoTweet(text)
}

func (cons *Consumer) Reply(text string, tweet *twitter.Tweet) error {
	return cons.apiClient.DoReply(text, tweet.Id)
}

func (cons *Consumer) Favorite(tweet *twitter.Tweet) error {
	return cons.apiClient.DoFavorite(tweet.Id)
}

func (cons *Consumer) Retweet(tweet *twitter.Tweet) error {
	return cons.apiClient.DoRetweet(tweet.Id)
}

func (cons *Consumer) GetScreenName() string {
	return cons.screenName
}
