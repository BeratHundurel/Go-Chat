package handlers

import (
	"go-chat/app/services"
	"go-chat/app/types"
	"go-chat/app/views/landing"
	"net/http"
)

func HandleLandingIndex(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("authentication")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	phone := cookie.Value

	user := services.GetUserByPhone(phone)
	if user.ID == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	var ChatViews []types.ChatView
	friends := services.GetFriendsByUser(user)
	for _, friend := range friends {
		lastMessage := services.GetLastMessage(user, friend)
		chatView := types.ChatView{
			User:        friend,
			LastMessage: lastMessage,
		}
		ChatViews = append(ChatViews, chatView)
	}
	availabe_users := services.GetAvailableUsers(user)

	view := types.UserView{
		AvailableUsers: availabe_users,
		CurrentUser:    user,
		ChatViews:      ChatViews,
	}

	landing.Index(view).Render(r.Context(), w)
}
