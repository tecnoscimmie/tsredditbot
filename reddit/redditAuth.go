package reddit

import (
	"encoding/json"
	"fmt"
	"github.com/tecnoscimmie/tsredditbot/support"
	"net/http"
	"net/url"
)

var requestTokenURL = "https://www.reddit.com/api/v1/access_token"
var baseOauthURL = "https://oauth.reddit.com/api/"
var defaultSubreddit = "tecnoscimmie"

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

// Session is a currently-active reddit session
type Session struct {
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
	Token        TokenResponse
}

// NewSession creates a Session
func NewSession(username string, password string, clientID string, clientSecret string) (Session, error) {
	s := Session{}
	s.Username = username
	s.Password = password
	s.ClientID = clientID
	s.ClientSecret = clientSecret
	err := s.RequestToken()

	if err != nil {
		return Session{}, err
	}

	return s, nil
}

// RequestToken requests a token for any given username, password, client ID and client secret.
// Does not return the TokenResponse, that will be saved inside the Session instead.
// Does return an error.
func (s *Session) RequestToken() error {

	client := &http.Client{}

	vals := url.Values{}
	vals.Set("grant_type", "password")
	vals.Set("username", s.Username)
	vals.Set("password", s.Password)

	req, err := createTokenRequest("POST", requestTokenURL, vals, s.ClientID, s.ClientSecret)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var token TokenResponse
	respDecoder := json.NewDecoder(resp.Body)
	err = respDecoder.Decode(&token)

	if err != nil {
		return err
	}

	s.Token = token

	return nil
}

// Post posts an URL to our subreddit
func (s *Session) Post(urlBeingPosted string) error {
	client := &http.Client{}

	pageTitle, err := support.GetPageTitle(urlBeingPosted)
	if err != nil {
		return err
	}

	vals := url.Values{}
	vals.Add("sr", defaultSubreddit)
	vals.Add("kind", "link")
	vals.Add("url", urlBeingPosted)
	vals.Add("title", pageTitle)

	req, err := createAuthenticatedRequest("POST", baseOauthURL+"submit", s.Token, vals)
	if err != nil {
		return err
	}

	_, err = client.Do(req) // _ -> resp

	if err != nil {
		return err
	}

	// need to parse resp and check for errors
	return nil
}

// Me pokes /api/v1/me
func (s *Session) Me() (Me, error) {
	client := &http.Client{}
	req, err := createAuthenticatedRequest("GET", baseOauthURL+"me", s.Token, url.Values{})
	if err != nil {
		return Me{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return Me{}, err
	}

	defer resp.Body.Close()
	var m Me
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&m)

	if err != nil {
		return Me{}, err
	}

	return m, nil
}
