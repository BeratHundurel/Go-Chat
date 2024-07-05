package app

import (
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"go-chat/app/handlers"
	"go-chat/app/middleware"
)

// Define your global middleware
func InitializeMiddleware(router *chi.Mux) {
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
}

func InitializeRoutes(router *chi.Mux) {
	router.Group(func(app chi.Router) {
		app.Get("/", handlers.HandleAuthIndex)
		app.With(middleware.AuthMiddleware).Get("/chat", handlers.HandleLandingIndex)
		app.Post("/register", handlers.HandleAuthRegister)
	})
}
