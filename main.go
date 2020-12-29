package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CookieHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cookies/{id}", CookieHandler)
	http.Handle("/", r)
}
