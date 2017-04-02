package twitter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/mrjones/oauth"
)

var (
	requestTokenURL   = "https://api.twitter.com/oauth/request_token"
	authorizeTokenURL = "https://api.twitter.com/oauth/authorize"
	accessTokenURL    = "https://api.twitter.com/oauth/access_token"
)

// Authenticate returns oauth access token.
func Authenticate() (string, string, error) {
	consumer := newTwitterConsumer()
	requestToken, url, err := consumer.GetRequestTokenAndUrl("")
	if err != nil {
		return "", "", err
	}

	openBrowser(url)
	pinCode := readPinCode()

	accessToken, err := consumer.AuthorizeToken(requestToken, pinCode)
	if err != nil {
		return "", "", err
	}

	return accessToken.Token, accessToken.Secret, nil
}

func newTwitterConsumer() *oauth.Consumer {
	provider := oauth.ServiceProvider{
		RequestTokenUrl:   requestTokenURL,
		AuthorizeTokenUrl: authorizeTokenURL,
		AccessTokenUrl:    accessTokenURL,
	}

	credential := NewCredential()

	return oauth.NewConsumer(
		credential.ConsumerKey,
		credential.ConsumerSecret,
		provider,
	)
}

func openBrowser(url string) {
	browser := "xdg-open"
	args := []string{url}
	if runtime.GOOS == "windows" {
		browser = "rundll32.exe"
		args = []string{"url.dll,FileProtocolHandler", url}
	} else if runtime.GOOS == "darwin" {
		browser = "open"
	} else if runtime.GOOS == "plan9" {
		browser = "plumb"
	}

	fmt.Println("Open this URL.")
	fmt.Println(url)
	browser, err := exec.LookPath(browser)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(browser, args...)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func readPinCode() string {
	fmt.Print("Enter PIN: ")
	stdin := bufio.NewReader(os.Stdin)
	input, err := stdin.ReadBytes('\n')
	if err != nil {
		log.Fatal("Canceled")
	}

	if input[len(input)-2] == '\r' {
		input = input[0 : len(input)-2]
	} else {
		input = input[0 : len(input)-1]
	}

	return string(input)
}
