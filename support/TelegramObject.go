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
