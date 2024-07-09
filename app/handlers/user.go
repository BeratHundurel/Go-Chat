package handlers

import (
	"go-chat/app/services"
	"go-chat/app/types"
	"net/http"
)

func FetchAvailableUsers(w http.ResponseWriter, r *http.Request) []types.User {
	cookie, err := r.Cookie("authentication")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil
	}
	phone := cookie.Value
	users := services.GetAvailableUsers(phone)
	return users
}

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
}
