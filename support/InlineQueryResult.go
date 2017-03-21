package support

// InlineQueryResultArticle represents a result for an inline request, for an URL
type InlineQueryResultArticle struct {
	Type                string              `json:"type"`
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	URL                 string              `json:"url"`
	HideURL             bool                `json:"hide_url"`
	InputMessageContent InputMessageContent `json:"input_message_content"`
}

// InputMessageContent is the real content of an InlineQueryResultArticle
type InputMessageContent struct {
	MessageText string `json:"message_text"`
}

// NewResultArticle builds and returns an InlineQueryResultArticle based on the parameters passed
func NewResultArticle(id string, url string, hideURL bool) InlineQueryResultArticle {
	var r InlineQueryResultArticle
	r.Type = "article"
	r.ID = id
	if title, err := GetPageTitle(url); err != nil {
		r.Title = "Not a valid URL :( Try with something like https://gsora.xyz"
	} else {
		r.Title = title
	}
	r.URL = url
	r.HideURL = hideURL
	r.InputMessageContent = InputMessageContent{MessageText: "Ready to be posted :D"}

	return r
}
