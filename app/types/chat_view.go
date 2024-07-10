package types

type ChatView struct {
	User        User    `json:"user"`
	LastMessage Message `json:"last_message"`
}
