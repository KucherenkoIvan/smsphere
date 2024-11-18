package http

import (
	"fmt"
	"log"
	"net/http"
	"smsphere/internal/http/controllers"
)

const (
	_ATTEMPTS = 25
	PORT      = 8080
)

var handlers map[string]http.HandlerFunc

func getRouteKey(method string, prefix string) string {
	return fmt.Sprintf("%s %s", method, prefix)
}

func route(method string, prefix string, handler http.HandlerFunc) {
	handlers[getRouteKey(method, prefix)] = handler

	http.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		key := getRouteKey(r.Method, prefix)
		handler := handlers[key]
		if handler == nil {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		} else {
			handler(w, r)
		}
	})
}

func Start() {
	handlers = make(map[string]http.HandlerFunc)

	route("GET", "/app/info", controllers.AppInfo)
	route("GET", "/api/app/counter", controllers.Counter)
	route("POST", "/api/app/auth/login", controllers.Login)
	route("POST", "/api/app/auth/register", controllers.Register)

	var error error
	for i := 0; i < _ATTEMPTS; i++ {
		address := fmt.Sprintf("%s%d", "localhost:", PORT+i)
		log.Printf("Starting http server on %s...", address)
		error = http.ListenAndServe(address, nil)
		log.Printf("Error %s", error)
	}

	log.Fatal(error)
}
