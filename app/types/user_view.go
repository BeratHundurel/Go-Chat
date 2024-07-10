package types

type UserView struct {
	AvailableUsers []User     `json:"available_users"`
	CurrentUser    User       `json:"current_user"`
	ChatViews      []ChatView `json:"chat_views"`
}
