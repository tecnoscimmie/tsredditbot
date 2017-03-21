package support

// ChosenInlineResult is what Telegram gives us when an user selects an inline result
type ChosenInlineResult struct {
	From struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
	} `json:"from"`
	InlineMessageID string `json:"inline_message_id"`
	Query           string `json:"query"`
	ResultID        string `json:"result_id"`
}
