package auth

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/big"
	"smsphere/internal/db"
	"smsphere/internal/db/models"
)

// TODO: move to env
const _SALT = "mmmhm salty"

func getPasswordHash(password string, pepper string) string {
	spicyPassword := fmt.Sprintf("%s%s%s", _SALT, password, pepper)

	hasher := crypto.SHA512.New()
	hasher.Write([]byte(spicyPassword))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func Login(login string, password string) bool {
	var candidate *models.AuthUser

	result := db.Connection.First(&candidate, models.AuthUser{Login: login})
	if result.Error != nil {
		log.Print(result.Error)

		return false
	}

	passwordHash := getPasswordHash(password, candidate.Pepper)

	if passwordHash == candidate.PasswordHash {
		return true
	}

	return false
}

func Register(login string, password string) (*models.AuthUser, error) {
	var candidate *models.AuthUser
	result := db.Connection.First(&candidate, models.AuthUser{Login: login})

	if result.Error == nil {
		return nil, errors.New("User already exists")
	} else {
		log.Println(result.Error)
	}

	pepper, error := rand.Int(rand.Reader, big.NewInt(1e9))
	if error != nil {
		return nil, error
	}

	passwordHash := getPasswordHash(password, pepper.String())

	record := models.AuthUser{Login: login, PasswordHash: passwordHash, Pepper: pepper.String()}

	result = db.Connection.Create(&record)
	if result.Error != nil {
		return nil, error
	}

	return &record, nil
}
