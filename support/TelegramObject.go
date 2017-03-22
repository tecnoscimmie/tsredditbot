package support

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// TelegramObject "Update" message model
type TelegramObject struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// DecodeJSON decodes some JSON into a TelegramObject
func (t *TelegramObject) DecodeJSON(r io.ReadCloser) error {
	d := json.NewDecoder(r)
	err := d.Decode(t)
	if err != nil {
		return err
	}

	return nil
}

// ReplyBackToChat replies to the chat referred by Message.Chat.ID
func (t *TelegramObject) ReplyBackToChat(c string) error {
	vals := url.Values{}
	vals.Set("chat_id", strconv.Itoa(t.Message.Chat.ID))
	vals.Set("text", c)

	_, err := http.PostForm(baseURL+"sendMessage", vals)
	if err != nil {
		return err
	}

	return nil
}
