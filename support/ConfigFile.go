package support

// ConfigFile represents the structure of a configuration file
type ConfigFile struct {
	CertPath string `json:"cert_path"`
	KeyPath  string `json:"key_path"`
	URL      string `json:"url"`
	Endpoint string `json:"endpoint"`
	BotToken string `json:"bot_token"`
	Port     string `json:"port"`
}
