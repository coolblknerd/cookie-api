package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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

	if rr.Body.String() != expected {
		t.Errorf("Response body didn't return expected value.\nGot: %v\nExpected: %v", rr.Body.String(), expected)
	}
}
