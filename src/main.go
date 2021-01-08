package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/coolblknerd/cookie-api/src/handler"
	"github.com/coolblknerd/cookie-api/src/helper"
	"github.com/gorilla/mux"
)

var configs = helper.SetUpConfigs()

// Why do I have duplicate documents in my database?
func main() {
	var wait time.Duration

	r := mux.NewRouter().StrictSlash(true)
	apiRouter := r.PathPrefix("/api").Subrouter() // /api will give access to all the API endpoints
	apiRouter.HandleFunc("/cookies", handler.GetCookieByID).Methods("GET")
	apiRouter.HandleFunc("/cookies", handler.GetCookies).Methods("GET")
	apiRouter.HandleFunc("/cookies", handler.CreateCookie).Methods("POST")
	apiRouter.HandleFunc("/cookies/{id}", handler.DeleteCookieByID).Methods("DELETE")
	apiRouter.HandleFunc("/cookies/{id}", handler.UpdateCookieByID).Methods("PUT")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         configs.Server.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	fmt.Printf("Listening on port %v\n", configs.Server.Port)

	// The following code allows for a graceful shutdown of the server
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
