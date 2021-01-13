package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/coolblknerd/cookie-api/src/helper"
	"github.com/coolblknerd/cookie-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

func cleanUpRecords() {
	collection := helper.ConnectDB("devCookies")

	result, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		log.Println("There was a problem cleaning up this record.")
	}

	log.Println("Removed records: ", result.DeletedCount)
}

// Not really understanding why this test fails when the expected output and the response body are the same.
func TestGetCookieByID(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/api/cookies?id=5fef7c0dd83dd77db9b8ec3b", nil)
	if err != nil {
		t.Error("The request returned a error")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCookieByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"_id":"5fef7c0dd83dd77db9b8ec3b","name":"chocolate-chip","size":"medium","ingredients":["chocolate chips","butter","sugar","dough"],"calories":"500 cals","location":"Crumbl Cookies","vegetarian":false}`

	if reflect.DeepEqual(rr.Body.String(), expected) == false {
		t.Errorf("Response body didn't return expected value.\nGot: %v\nExpected: %v", rr.Body.String(), expected)
	}
}

func TestCreateCookie(t *testing.T) {
	newCookie := models.Cookie{
		Name: "Hazelnut Fudge",
		Size: "Large",
		Ingredients: []string{
			"butter",
			"hazelnut",
			"fudge",
		},
		Calories:   "400 cals",
		Location:   "Katy, TX",
		Vegetarian: false,
	}

	reqBytes, _ := json.Marshal(newCookie)
	reqBody := bytes.NewReader(reqBytes)
	req, err := http.NewRequest("POST", "localhost:8080/api/cookies", reqBody)
	if err != nil {
		t.Error("The request returned a error")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCookie)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	cleanUpRecords()
}
