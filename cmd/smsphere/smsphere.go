package main

import (
	"fmt"
	"smsphere/internal/auth"
	"smsphere/internal/db"
	"smsphere/internal/db/models"
	http_server "smsphere/internal/http"
)

func main() {
	db.Connect()

	db.Connection.AutoMigrate(&models.Log{})
	db.Connection.AutoMigrate(&models.AuthUser{})

	db.Connection.Create(&models.Log{Text: "App started"})

	// create
	log := models.Log{Text: "This log will be deleted"}
	result := db.Connection.Create(&log)

	// Полезное
	// result.Error
	// result.RowsAffected
	fmt.Println(result.RowsAffected)

	// read
	var record models.Log
	db.Connection.First(&record, log.ID)

	// update
	record.Text = "Text has been updated"
	db.Connection.Save(&record)

	// delete
	db.Connection.Delete(&record)

	aboba, error := auth.Register("aboba_4411", "very_strong_password")
	login, error := auth.Login("aboba_4411", "very_strong_password")
	fmt.Println("\n\n\nregister result\n\n\n")
	fmt.Println(aboba)
	fmt.Println(error)
	fmt.Println("\n\n\nlogin result\n\n\n")
	fmt.Println(login)
	fmt.Println(error)
	fmt.Println("\n\n\nend login result\n\n\n")

	// init http router and start listening
	http_server.Start()
}
