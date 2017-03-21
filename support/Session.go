package support

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL string

// Session is a Telegram bot session
type Session struct {
	Token         string
	Port          string
	Endpoint      string
	URL           string
	Configuration ConfigFile
}

// NewSession returns a new Telegram bot session
func NewSession() (Session, error) {
	conf, err := CheckConfigFile()
	if err != nil {
		return Session{}, err
	}

	var s Session
	s.Token = conf.BotToken
	s.Port = conf.Port
	s.Endpoint = conf.Endpoint
	s.URL = conf.URL
	s.Configuration = conf

	baseURL = "https://api.telegram.org/bot" + s.Token + "/"

	// setup webhook already
	err = s.SetupWebHook()
	if err != nil {
		return Session{}, err
	}

	return s, nil
}

// SetupWebHook setup the Telegram webhook for the running server
func (s *Session) SetupWebHook() error {
	// TODO: handle webhook errors
	_, err := http.PostForm(baseURL+"setWebhook", url.Values{"url": {s.URL + ":" + s.Port + "/" + s.Endpoint}})
	if err != nil {
		return err
	}

	return nil
}

// PrintBotInformations prints some informations about the Bot the session represents
func (s *Session) PrintBotInformations() error {
	data, err := http.Get(baseURL + "getMe")
	if err != nil {
		return err
	}
	var u BotInfo
	err = u.DecodeJSON(data.Body)
	if err != nil {
		return err
	}

	fmt.Println("Bot handle: @" + u.Result.Username)
	fmt.Println("Bot name: " + u.Result.FirstName)

	return err
}

// ReplyToInlineQuery replies to the inline query contained into the TelegramObject we're referencing
func (s *Session) ReplyToInlineQuery(t TelegramObject) error {
	article := []InlineQueryResultArticle{NewResultArticle(t.InlineQuery.ID, t.InlineQuery.Query, false)}

	enc, err := json.Marshal(article)

	if err != nil {
		return err
	}

	v := url.Values{}
	v.Add("inline_query_id", t.InlineQuery.ID)
	v.Add("results", string(enc))

	_, err = http.PostForm(baseURL+"answerInlineQuery", v)
	if err != nil {
		return err
	}

	return nil
}
