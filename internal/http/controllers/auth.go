package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"smsphere/internal/auth"
	"smsphere/internal/db/models"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: move to env
const _TOKEN_KEY = "jlkwjlkjanncnawld"

type Request struct {
	Login    string
	Password string
}

type Response struct {
	Token string
}

func generateJWT(user *models.AuthUser) (string, error) {
	tokenGen := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": user.Login,
		"id":    user.ID,
	})

	token, error := tokenGen.SignedString([]byte(_TOKEN_KEY))
	if error != nil {
		return "", error
	}

	return token, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body Request

	error := decoder.Decode(&body)
	if error != nil {
		log.Println(error)

		http.Error(w, "Bad request error", http.StatusBadRequest)
		return
	}

	user, error := auth.Login(body.Login, body.Password)
	if error != nil {
		http.Error(w, "Login or password incorrect", http.StatusBadRequest)
		return
	}

	token, error := generateJWT(user)
	if error != nil {
		log.Println(error)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Token: token})
}

func Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body Request

	error := decoder.Decode(&body)
	if error != nil {
		panic(error)
	}

	user, error := auth.Register(body.Login, body.Password)
	if error != nil {
		http.Error(w, "Registration error", http.StatusBadRequest)
	}

	token, error := generateJWT(user)
	if error != nil {
		log.Println(error)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Token: token})
}
