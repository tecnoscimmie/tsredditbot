package support

// InlineQueryReply is a reply ready to be sent to telegram
type InlineQueryReply struct {
	InlineQueryID string                     `json:"inline_query_id"`
	Results       []InlineQueryResultArticle `json:"results"`
}
