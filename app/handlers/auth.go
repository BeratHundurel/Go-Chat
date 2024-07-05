package handlers

import (
	"go-chat/app/services"
	"go-chat/app/types"
	"go-chat/app/views/auth"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func HandleAuthIndex(w http.ResponseWriter, r *http.Request) {
	auth.Index(auth.RegisterFormValues{}).Render(r.Context(), w)
}

func HandleAuthRegister(w http.ResponseWriter, r *http.Request) {
	var form auth.RegisterFormValues
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := decoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usernameExists := services.CheckUsername(form.Username)
	if usernameExists {
		http.Error(w, "Username is already taken", http.StatusBadRequest)
		return
	}
	
	user := types.User{
		Username: form.Username,
		Phone:    form.Phone,
		Password: form.Password,
	}
	services.RegisterUser(user)

	http.SetCookie(w,&http.Cookie{
		Name: "authentication",
		Value: user.Username,
		Expires: time.Now().Add(128 * time.Hour),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/chat", http.StatusFound)
}