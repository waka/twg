package twitter

// Credential is app settings.
type Credential struct {
	ConsumerKey    string
	ConsumerSecret string
}

// NewCredential return this application credential info.
func NewCredential() *Credential {
	return &Credential{
		ConsumerKey:    "MYTtfkwR1SdJsIGvbMay5KtC9",
		ConsumerSecret: "lbRXrUpP9QxN46bCXRRLBtXiNphtNmfNDfH25vD99G79nrgj7p",
	}
}
