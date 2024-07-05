package handlers

import (
	"go-chat/app/views/landing"
	"net/http"
)

func HandleLandingIndex(w http.ResponseWriter, r *http.Request) {
	landing.Index().Render(r.Context(), w)
}
