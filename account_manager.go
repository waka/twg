package main

import (
	"errors"

	"github.com/waka/twg/twitter"
)

// AccountManager is useful for API.
type AccountManager struct {
	config    *Config
	apiClient *twitter.Client
}

// NewAccountManager returns instance.
func NewAccountManager() *AccountManager {
	return &AccountManager{config: LoadConfig()}
}

// GetAPIClient returns twitter api client.
func (am *AccountManager) GetAPIClient() *twitter.Client {
	return am.apiClient
}

// Auth is authentication method.
func (am *AccountManager) Auth() error {
	if am.config.IsAuthenticated() == false {
		token, secret, err := twitter.Authenticate()
		if err != nil {
			return err
		}

		am.config.SetAccessToken(token, secret)
		if am.config.Save() == false {
			return errors.New("Failed to save config")
		}
	}

	// set api client
	am.apiClient = twitter.NewClient(
		am.config.AccessToken,
		am.config.AccessTokenSecret,
	)

	return nil
}
