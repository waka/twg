package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

// Config has api access token.
type Config struct {
	AccessToken       string
	AccessTokenSecret string
	ScreenName        string
}

// LoadConfig load json from local.
func LoadConfig() *Config {
	config := &Config{}
	path := configFilePath()

	if fileExists(path) {
		data, err := ioutil.ReadFile(path)
		if err == nil {
			decoder := json.NewDecoder(bytes.NewReader(data))
			decoder.Decode(config)
		}
	}

	return config
}

// Save config to local
func (config *Config) Save() bool {
	json, err := json.Marshal(config)
	if err != nil {
		return false
	}
	ioutil.WriteFile(configFilePath(), json, 0644)

	return true
}

// IsAuthenticated return whether you authenticated?
func (config *Config) IsAuthenticated() bool {
	return config.AccessToken != "" && config.AccessTokenSecret != ""
}

// SetAccessToken set token to config.
func (config *Config) Set(token string, secret string, screenName string) {
	config.AccessToken = token
	config.AccessTokenSecret = secret
	config.ScreenName = screenName
}

func fileExists(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	return os.IsNotExist(err) == false
}

func configFilePath() string {
	target, err := user.Current()
	if err != nil {
		return ""
	}
	return target.HomeDir + "/.twg"
}
