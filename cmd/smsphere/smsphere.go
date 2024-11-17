package main

import (
	"fmt"
	"smsphere/internal/db"
	"smsphere/internal/db/models"
	router "smsphere/internal/http"
)

func main() {
	db.Connect()

	db.Connection.AutoMigrate(&models.Log{})

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

	// init http router and start listening
	router.InitRouter()
}
