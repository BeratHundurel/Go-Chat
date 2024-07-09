package handlers

import (
	"go-chat/app/views/landing"
	"net/http"
)

func HandleLandingIndex(w http.ResponseWriter, r *http.Request) {
	availabe_users := FetchAvailableUsers(w, r)
	landing.Index(availabe_users).Render(r.Context(), w)
}
