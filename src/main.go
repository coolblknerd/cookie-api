package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	c "./config"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

func GetCookie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	collection := client.Database(configuration.Database.Name).Collection("cookies")
	res, err := collection.
		fmt.Fprintf(w, "Cookie ID: %v", id)
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {

}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {

}

func UpdateCookie(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// -- Set-up Configurations --

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration c.Configurations
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// -- Set-up Database Client --

	client, err := mongo.NewClient(options.Client().ApplyURI(configuration.Database.URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mongo Successfully connected.")
	defer client.Disconnect(ctx)

	// -- Set-up Server

	r := mux.NewRouter()
	r.HandleFunc("/api/cookies/{id}", GetCookie).Methods("GET")
	r.HandleFunc("/api/cookies/{id}", CreateCookie).Methods("POST")
	r.HandleFunc("/api/cookies/{id}", DeleteCookie).Methods("DELETE")
	r.HandleFunc("/api/cookies/{id}", UpdateCookie).Methods("PUT")

	http.Handle("/", r)
	http.ListenAndServe(configuration.Server.Port, r)
}
