package support

// BotInfo contains informations about the bot
type BotInfo struct {
	Ok     bool `json:"ok"`
	Result User `json:"result"`
}
