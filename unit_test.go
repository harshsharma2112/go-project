package main

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"
)

func TestHealthHandler(t *testing.T) {
	// 1.creating a fake http request for my func
	t.Log("working till here")
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Test is failed due to %v", err) //fatal is used becoz we wanted to stop the test here only

	}
	//2.recording my response
	rr := httptest.NewRecorder()
	//3.calling my healthhandler function
	//http.HandlerFunc is a type adapter in the Go net/http package.
	//It converts your ordinary function func(w http.ResponseWriter, r *http.Request) into an object that implements the http.Handler interface.
	//logic written in white notebook for refernece
	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong staus got %v but want %v", status, http.StatusOK)
		//errorf to support formatting
	}

	//5. to check the json Part
	expected := `"status":"UP"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want to contain %v", rr.Body.String(), expected)
	}
}

func TestRoothandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("fail rqst error %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Roothandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong staus got %v but want %v", status, http.StatusOK)
		//errorf to support formatting
	}

	expected := "Aplication is running and up"
	if rr.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())

	}
}
