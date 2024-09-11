package chat

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	To      string `json:"to"`
}
