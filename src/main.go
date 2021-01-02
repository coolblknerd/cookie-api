package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/coolblknerd/cookie-api/src/helper"
	"github.com/coolblknerd/cookie-api/src/models"
	"github.com/gorilla/mux"
)

var collection = helper.ConnectDB()
var configs = helper.SetUpConfigs()

func GetCookies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cookies []*models.Cookie

	findOptions := options.Find()
	findOptions.SetLimit(10)

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var cookie models.Cookie
		err := cur.Decode(&cookie)
		if err != nil {
			log.Fatal(err)
		}

		cookies = append(cookies, &cookie)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", cookies)

}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cookie models.Cookie
	vars := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&cookie)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(cookie)
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cookie models.Cookie
	err := json.NewDecoder(r.Body).Decode(&cookie)

	insertResult, err := collection.InsertOne(context.TODO(), cookie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func UpdateCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cookie models.Cookie
	err := json.NewDecoder(r.Body).Decode(&cookie)

	vars := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"name":        cookie.Name,
			"size":        cookie.Size,
			"ingredients": cookie.Ingredients,
			"calories":    cookie.Calories,
			"location":    cookie.Location,
			"vegetarian":  cookie.Vegetarian,
		},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// Why do I have duplicate documents in my database?
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/cookies/{id}", GetCookie).Methods("GET")
	r.HandleFunc("/api/cookies", GetCookies).Methods("GET")
	r.HandleFunc("/api/cookies", CreateCookie).Methods("POST")
	r.HandleFunc("/api/cookies/{id}", DeleteCookie).Methods("DELETE")
	r.HandleFunc("/api/cookies/{id}", UpdateCookie).Methods("PUT")

	http.Handle("/", r)
	http.ListenAndServe(configs.Server.Port, r)
}
