package main

import (
	"errors"

	"github.com/waka/twg/twitter"
)

type AccountManager struct {
	config    *Config
	apiClient *twitter.Client
}

func NewAccountManager() *AccountManager {
	return &AccountManager{config: LoadConfig()}
}

func (self *AccountManager) GetApiClient() *twitter.Client {
	return self.apiClient
}

func (self *AccountManager) Auth() error {
	if self.config.IsAuthenticated() == false {
		token, secret, err := twitter.Authenticate()
		if err != nil {
			return err
		}

		self.config.SetAccessToken(token, secret)
		if self.config.Save() == false {
			return errors.New("Failed to save config")
		}
	}

	// set api client
	self.apiClient = twitter.NewClient(
		self.config.AccessToken,
		self.config.AccessTokenSecret,
	)

	return nil
}
