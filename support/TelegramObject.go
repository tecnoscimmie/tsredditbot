package support

// TelegramObject "Update" message model
type TelegramObject struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}
