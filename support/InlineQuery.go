package support

// InlineQuery is what Telegram gives use when an user make an Inline query
type InlineQuery struct {
	ID   string `json:"id"`
	From struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
	} `json:"from"`
	Query  string `json:"query"`
	Offset string `json:"offset"`
}
