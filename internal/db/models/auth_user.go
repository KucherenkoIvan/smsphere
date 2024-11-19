package models

import "gorm.io/gorm"

type AuthUser struct {
	gorm.Model

	Login        string
	Pepper       string
	PasswordHash string
}
