package main

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"
)

//✅ Unit Tests Are Not Meant to Never Fail
//Unit tests are designed to fail when your function does not behave as expected.

func TestHealthHandler(t *testing.T) {
	// 1.creating a fake http request for my func
	t.Log("working till here")
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Test is failed due to %v", err) //fatal is used becoz we wanted to stop the test here only

	}
	//2.recording my response'
	// type ResponseRecorder struct {
	// Code      int           // the HTTP status code (e.g., 200, 404, etc.)
	// HeaderMap http.Header   // headers sent by the handler
	// Body      *bytes.Buffer // the body written by the handler
	// Flushed   bool          // whether Flush was called
	// // ... some internal fields
	rr := httptest.NewRecorder()

	//3.calling my healthhandler function
	//http.HandlerFunc is a type adapter in the Go net/http package.
	//It converts your ordinary function func(w http.ResponseWriter, r *http.Request) into an object that implements the http.Handler interface.
	//logic written in white notebook for refernece
	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(rr, req) // test will fail only after the execution of this line

	//In Go, http.StatusOK is just a constant defined in the net/http package. Internally, it's just an integer:

	if status := rr.Code; status != http.StatusOK {  //; is supported due to if conditition
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

func TestDetailhandler(t *testing.T) {

	req, err := http.NewRequest("GET", "deatils", nil)
	if err != nil {
		t.Errorf("following error has occured%v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Detailhandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong staus got %v but want %v", status, http.StatusOK) //`` need to look into this code again
		//errorf to support formatting
	}

    body := rr.Body.String()
	if !strings.Contains(body, "hostname") || !strings.Contains(body, "IP") {
		t.Errorf("Response body missing expected keys: %s", body)
	}
}
