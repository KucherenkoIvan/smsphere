package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"smsphere/internal/auth"
)

// TODO: move to env
const _TOKEN_KEY = "jlkwjlkjanncnawld"

type RequestBody struct {
	Login    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body RequestBody

	error := decoder.Decode(&body)
	if error != nil {
		log.Println(error)

		http.Error(w, "Bad request error", http.StatusBadRequest)
	}

	login := auth.Login(body.Login, body.Password)

	if !login {
		http.Error(w, "Login or password incorrect", http.StatusBadRequest)
	}

	// TODO: return jwt
}

func Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body RequestBody

	error := decoder.Decode(&body)
	if error != nil {
		panic(error)
	}

	user, error := auth.Register(body.Login, body.Password)
	if error != nil {
		http.Error(w, "Registration error", http.StatusBadRequest)
	}

	// TODO: return jwt
	log.Println(user)
}
