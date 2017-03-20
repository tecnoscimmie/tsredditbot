package support

import (
	"encoding/json"
	"io"
)

// BotInfo contains informations about the bot
type BotInfo struct {
	Ok     bool `json:"ok"`
	Result User `json:"result"`
}

// DecodeJSON decodes some JSON into a BotInfo
func (u *BotInfo) DecodeJSON(r io.ReadCloser) error {
	d := json.NewDecoder(r)
	err := d.Decode(u)
	if err != nil {
		return err
	}

	return nil
}
