package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var requestTokenURL = "https://www.reddit.com/api/v1/access_token"
var baseOauthURL = "https://oauth.reddit.com/api/v1/"

// TokenResponse is what reddit replies us with when requesting a token using RequestToken()
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GetAuthString returns a well-formed authentication string
func (t *TokenResponse) GetAuthString() string {
	return fmt.Sprintf("%s %s", t.TokenType, t.AccessToken)
}

// RequestToken requests a token for any given username, password, client ID and client secret
func RequestToken(username string, password string, clientID string, clientSecret string) (TokenResponse, error) {

	client := &http.Client{}

	vals := url.Values{}
	vals.Set("grant_type", "password")
	vals.Set("username", username)
	vals.Set("password", password)

	req, err := createTokenRequest("POST", requestTokenURL, vals, clientID, clientSecret)

	resp, err := client.Do(req)
	if err != nil {
		return TokenResponse{}, err
	}
	defer resp.Body.Close()

	var token TokenResponse
	respDecoder := json.NewDecoder(resp.Body)
	err = respDecoder.Decode(&token)

	if err != nil {
		return TokenResponse{}, err
	}

	return token, nil
}

func Me(t TokenResponse) (string, error) {
	client := &http.Client{}
	req, err := createAuthenticatedRequest("GET", baseOauthURL+"me", t)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b), nil
}
