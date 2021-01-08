package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/coolblknerd/cookie-api/src/helper"
	"github.com/coolblknerd/cookie-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	if os.Getenv("COOKIE_ENV") != "prod" {
		collection = helper.ConnectDB("dev-cookies")
		log.Println("Connected to dev environment")
	}
}

// GetCookies : Get all Cookies
// URL : /cookies
// Parameters: none
// Method: GET
// Output: JSON Encoded Entries object if found else JSON Encoded Exception.
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

	err = json.NewEncoder(w).Encode(cookies)
	if err != nil {
		log.Fatalf("There was a problem returning the response in json: %v", err)
	}
}

// GetCookieByID : Get Cookie by ID
// URL : /cookies/id=?
// Parameters: none
// Method: GET
// Output: JSON Encoded Entries object if found else JSON Encoded Exception.
func GetCookieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cookie models.Cookie

	query := r.URL.Query().Get("id")

	id, _ := primitive.ObjectIDFromHex(query)
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&cookie)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(cookie)
	if err != nil {
		log.Fatalf("There was a problem returning the response in json: %v", err)
	}
}

// CreateCookie : Create a Cookie
// URL : /cookies
// Parameters: {
// 		"name": Sting,
// 		"size": String,
//     	"ingredients": []String
//     	"calories": String,
//     	"location": String,
//     	"vegetarian": Bool
// }
// Method: POST
// Output: Returns a successful status code
func CreateCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cookie models.Cookie
	err := json.NewDecoder(r.Body).Decode(&cookie)
	if err != nil {
		log.Fatalf("There was a problem returning the response in json: %v", err)
	}

	insertResult, err := collection.InsertOne(context.TODO(), cookie)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
}

// DeleteCookieByID : Delete Cookie by ID
// URL : /cookies/id=?
// Parameters: none
// Method: GET
// Output: JSON Encoded Entries object if found else JSON Encoded Exception.
func DeleteCookieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("id")

	id, _ := primitive.ObjectIDFromHex(query)
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

// UpdateCookieByID : Update a Cookie by ID
// URL : /cookies?id=?
// Parameters: {
// 		"name": Sting,
// 		"size": String,
//     	"ingredients": []String
//     	"calories": String,
//     	"location": String,
//     	"vegetarian": Bool
// }
// Method: PUT
// Output: Returns a successful status code
func UpdateCookieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cookie models.Cookie
	err := json.NewDecoder(r.Body).Decode(&cookie)
	if err != nil {
		log.Fatalf("There was a problem reading the response: %v", err)
	}
	query := r.URL.Query().Get("id")

	id, _ := primitive.ObjectIDFromHex(query)
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

	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
