package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// 1.creating a fake http request for my func
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal("Test is failed due to %v", err) //fatal is used becoz we wanted to stop the test here only

	}
	//2.recording my response
	rr := httptest.NewRecorder()
	//3.calling my healthhandler function
	//http.HandlerFunc is a type adapter in the Go net/http package.
	//It converts your ordinary function func(w http.ResponseWriter, r *http.Request) into an object that implements the http.Handler interface.
	//logic written in white notebook for refernece
	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(rr, req)

	
	status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 5. Optionally, check if the response body contains the expected JSON
	expected := `"status":"UP"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want to contain %v",
			rr.Body.String(), expected)
	}

}
