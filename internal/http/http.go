package http

import (
	"fmt"
	"log"
	"net/http"
	"smsphere/internal/http/controllers"
)

const (
	ATTEMPTS = 5
	PORT     = 8080
)

func InitRouter() {
	http.HandleFunc("/app/info", controllers.AppInfo)
	http.HandleFunc("/api/app/counter", controllers.Counter)

	var error error
	for i := 0; i < ATTEMPTS; i++ {
		address := fmt.Sprintf("%s%d", ":", PORT+i)
		error = http.ListenAndServe(address, nil)
	}

	log.Fatal(error)
}
