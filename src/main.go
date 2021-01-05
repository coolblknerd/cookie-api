package main

import (
	"log"
	"net/http"

	"github.com/coolblknerd/cookie-api/src/handler"
	"github.com/coolblknerd/cookie-api/src/helper"
	"github.com/gorilla/mux"
)

var configs = helper.SetUpConfigs()

// Why do I have duplicate documents in my database?
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/cookies/{id}", handler.GetCookie).Methods("GET")
	r.HandleFunc("/api/cookies", handler.GetCookies).Methods("GET")
	r.HandleFunc("/api/cookies", handler.CreateCookie).Methods("POST")
	r.HandleFunc("/api/cookies/{id}", handler.DeleteCookie).Methods("DELETE")
	r.HandleFunc("/api/cookies/{id}", handler.UpdateCookie).Methods("PUT")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(configs.Server.Port, r))
}
