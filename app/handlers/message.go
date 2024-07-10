package handlers

import (
	"go-chat/app/helpers"
	"go-chat/app/services"
	"go-chat/app/types"
	"go-chat/app/views/landing"
	"net/http"
)

func HandleChat(w http.ResponseWriter, r *http.Request) {
	// Handle the chat
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}
	friendId := r.FormValue("id")
	friendID, err := helpers.ReturnIdAsIntFromString(friendId)
	if err != nil{
		http.Error(w, "Failed to parse friend id", http.StatusInternalServerError)
		return
	}

	chat, err := services.GetMessagesByFriendId(friendID)
	if err != nil {
		http.Error(w, "Failed to get chat", http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("authentication")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	phone := cookie.Value
	user := services.GetUserByPhone(phone)
	receiver := services.GetById(friendID)

	view := types.MessageView{
		Sender:   user,
		Messages: chat,
		Receiver: receiver,
	}

	// Render the chat
	landing.Messages(view).Render(r.Context(), w)
}
