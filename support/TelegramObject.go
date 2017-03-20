package support

import (
	"encoding/json"
	"io"
)

// TelegramObject "Update" message model
type TelegramObject struct {
	UpdateID           int                `json:"update_id"`
	Message            Message            `json:"message"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	InlineQuery        InlineQuery        `json:"inline_query"`
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

// HasInlineQuery checks if its TelegramObject has Inline query data
func (t *TelegramObject) HasInlineQuery() bool {
	if (InlineQuery{}) != t.InlineQuery || (ChosenInlineResult{}) != t.ChosenInlineResult {
		return false
	}

	return true
}
