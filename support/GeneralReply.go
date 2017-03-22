package support

// GeneralReply is any general telegram reply
type GeneralReply struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Result      Chat   `json:"result"`
}
