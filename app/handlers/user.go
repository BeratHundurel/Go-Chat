package handlers

import (
	"go-chat/app/services"
	"go-chat/app/views/components"
	"net/http"
)

func HandleAddFriendToUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("authentication")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	phone := cookie.Value
	user := services.GetUserByPhone(phone)

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}
	friendId := r.FormValue("id")
	services.AddFriend(user, friendId)

	available_users := services.GetAvailableUsers(user)
	components.FriendList(available_users).Render(r.Context(), w)
}
