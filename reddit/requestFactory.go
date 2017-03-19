package reddit

import (
	"net/http"
	"net/url"
	"strings"
)

var userAgent = "tsredditbot/0.1 by gsora (aka peppelakappa)"

func createTokenRequest(reqType string, url string, body url.Values, username string, password string) (*http.Request, error) {
	req, err := http.NewRequest(reqType, url, strings.NewReader(body.Encode()))
	if err != nil {
		return &http.Request{}, err
	}

	req.Header.Add("User-Agent", userAgent)
	req.SetBasicAuth(username, password)

	return req, nil
}

func createAuthenticatedRequest(reqType string, url string, token TokenResponse, body url.Values) (*http.Request, error) {
	req, err := http.NewRequest(reqType, url, strings.NewReader(body.Encode()))
	if err != nil {
		return &http.Request{}, err
	}
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Authorization", token.GetAuthString())
	return req, nil
}
