package controllers

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	mux   = new(sync.Mutex)
	count = 0
)

func Counter(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	count++
	mux.Unlock()
	fmt.Fprintln(w, count)
}
