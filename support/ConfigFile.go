package support

import (
	"bytes"
	"encoding/json"
	"io"
)

// ConfigFile represents the structure of a configuration file
type ConfigFile struct {
	CertPath    string `json:"cert_path"`
	KeyPath     string `json:"key_path"`
	URL         string `json:"url"`
	Endpoint    string `json:"endpoint"`
	BotToken    string `json:"bot_token"`
	Port        string `json:"port"`
	GroupHandle string `json:"group_handle"`
}

// DecodeConfigFile decodes a JSON configuration file into a ConfigFile
func (c *ConfigFile) DecodeConfigFile(r io.Reader) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	d := json.NewDecoder(buf)
	err := d.Decode(c)
	if err != nil {
		return err
	}

	return nil
}
