package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coolblknerd/cookie-api/src/handler"
	"github.com/coolblknerd/cookie-api/src/helper"
	"github.com/gorilla/mux"
)

var configs = helper.SetUpConfigs()

// Why do I have duplicate documents in my database?
func main() {
	r := mux.NewRouter().StrictSlash(true)
	apiRouter := r.PathPrefix("/api").Subrouter() // /api will give access to all the API endpoints
	apiRouter.HandleFunc("/cookies", handler.GetCookieByID).Methods("GET")
	apiRouter.HandleFunc("/cookies", handler.GetCookies).Methods("GET")
	apiRouter.HandleFunc("/cookies", handler.CreateCookie).Methods("POST")
	apiRouter.HandleFunc("/cookies/{id}", handler.DeleteCookieByID).Methods("DELETE")
	apiRouter.HandleFunc("/cookies/{id}", handler.UpdateCookieByID).Methods("PUT")
	http.Handle("/", r)
	fmt.Printf("Listening on port %v", configs.Server.Port)
	log.Fatal(http.ListenAndServe(configs.Server.Port, r))
}
