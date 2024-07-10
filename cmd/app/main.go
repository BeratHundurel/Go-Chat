package main

import (
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"go-chat/app"
	"go-chat/app/assets"
	"go-chat/app/db"
	"go-chat/app/hub"
	"go-chat/app/types"
	"log"
	"net/http"
	"os"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	serveStaticFiles(router)
	app.InitializeRoutes(router)
	listenAddr := os.Getenv("HTTP_LISTEN_ADDR")
	hub := hub.ReturnHub()
	go hub.Run()
	http.ListenAndServe(listenAddr, router)
}

func serveStaticFiles(router *chi.Mux) {
	if isDevelopment() {
		router.Handle("/public/*", disableCache(staticDev("public")))
		router.Handle("/app/assets/*", disableCache(staticDev("app/assets")))
	} else if isProduction() {
		router.Handle("/public/*", staticProd("public"))
		router.Handle("/app/assets/*", staticProd("app/assets"))
	}
}

func staticDev(directory string) http.Handler {
	return http.StripPrefix("/"+directory+"/", http.FileServerFS(os.DirFS(directory)))
}

func staticProd(directory string) http.Handler {
	return http.StripPrefix("/"+directory+"/", http.FileServerFS(assets.AssetsFS))
}

func isDevelopment() bool {
	return os.Getenv("ENV") == "development"
}

func isProduction() bool {
	return os.Getenv("ENV") == "production"
}

func disableCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func Migrate() {
	err := db.Get().AutoMigrate(&types.User{}, &types.Message{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		log.Println("Migration successful")
	}
}
