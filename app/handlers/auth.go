package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"go-chat/app/services"
	"go-chat/app/views/auth"
	"net/http"
	"time"
)

var decoder = schema.NewDecoder()
var validate = validator.New()

func HandleLoginGET(w http.ResponseWriter, r *http.Request) {
	auth.Index().Render(r.Context(), w)
}

func HandleLoginPOST(w http.ResponseWriter, r *http.Request) {
	var form auth.LoginFormValues
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := decoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := services.FindUserByUsername(form.Username)
	if user == nil || user.Password != form.Password {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "authentication",
		Value:    user.Phone,
		Expires:  time.Now().Add(128 * time.Hour),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleAuthRegisterGET(w http.ResponseWriter, r *http.Request) {
	auth.Register().Render(r.Context(), w)
}

func HandleAuthRegisterPOST(w http.ResponseWriter, r *http.Request) {
	var form auth.RegisterFormValues
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := decoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usernameExists := services.CheckUsername(form.Username)
	if usernameExists {
		http.Error(w, "Username is already taken", http.StatusBadRequest)
		return
	}

	services.RegisterUser(form)

	http.SetCookie(w, &http.Cookie{
		Name:     "authentication",
		Value:    form.Phone,
		Expires:  time.Now().Add(128 * time.Hour),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
