package support

import (
	"bytes"
	"encoding/json"
	"io"
)

// LoadJSONToTelegramObject loads json resulting from an HTTP request and
// returns a TelegramObject
func LoadJSONToTelegramObject(r io.ReadCloser) TelegramObject {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	var m TelegramObject
	json.Unmarshal(buf.Bytes(), &m)

	return m
}

// LoadJSONToConfigFile loads the configuration file (written in json) and
// returns a ConfigFile
func LoadJSONToConfigFile(r io.Reader) (ConfigFile, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	var m ConfigFile
	err := json.Unmarshal(buf.Bytes(), &m)

	return m, err
}

// LoadJSONToUser loads informations about the bot, provided from telegram, and
// returns a User
func LoadJSONToUser(r io.ReadCloser) (User, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	var m BotInfo
	err := json.Unmarshal(buf.Bytes(), &m)
	return m.Result, err
}
