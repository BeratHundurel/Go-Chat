package app

import (
	"go-chat/app/handlers"
	"go-chat/app/hub"
	"go-chat/app/middleware"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

// Define your global middleware
func InitializeMiddleware(router *chi.Mux) {
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
}

func InitializeRoutes(router *chi.Mux) {
	router.Group(func(app chi.Router) {
		app.With(middleware.AuthMiddleware).Get("/", handlers.HandleLandingIndex)
		app.Get("/login", handlers.HandleLoginGET)
		app.Get("/register", handlers.HandleAuthRegisterGET)
		app.Get("/ws", hub.HandleWebSocket)
		app.Post("/chat", handlers.HandleChat)
		app.Post("/login", handlers.HandleLoginPOST)
		app.Post("/register", handlers.HandleAuthRegisterPOST)
		app.Post("/add-friends", handlers.HandleAddFriendToUser)
	})
}
