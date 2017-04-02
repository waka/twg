package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

type Config struct {
	AccessToken       string
	AccessTokenSecret string
}

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

func (self *Config) Save() bool {
	json, err := json.Marshal(self)
	if err != nil {
		return false
	}
	ioutil.WriteFile(configFilePath(), json, 0644)

	return true
}

func (self *Config) IsAuthenticated() bool {
	return self.AccessToken != "" && self.AccessTokenSecret != ""
}

func (self *Config) SetAccessToken(token string, secret string) {
	self.AccessToken = token
	self.AccessTokenSecret = secret
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
