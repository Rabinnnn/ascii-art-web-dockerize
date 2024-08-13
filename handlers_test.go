package main

import (
	"ascii-art-web/handlers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

// Test funcionality of handler function Index
func TestIndex(t *testing.T) {

	testCases := []struct {
		name       string //test case name
		method     string
		path       string
		expected   string
		statusCode int
	}{
		{"GET /", http.MethodGet, "/", "templates/index.html", http.StatusMovedPermanently},
		{"GET /home", http.MethodGet, "/home", "templates/index.html", http.StatusOK},
		{"GET /aboutus", http.MethodGet, "/aboutus", "templates/aboutus.html", http.StatusOK},
		{"GET /404", http.MethodGet, "/404", "templates/404.html", http.StatusOK},


	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { //subtest for each case with the specified name
			req, err := http.NewRequest(tc.method, tc.path, nil) //mock http request with empty body
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder() //capture response
			handler := http.HandlerFunc(handlers.Index)
			handler.ServeHTTP(rr, req) //call handler with mock request and response from recorder

			if status := rr.Code; status != tc.statusCode { //check if recorder status code matches
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.statusCode)
			}

		})
	}
}

// Test functionality of handler function ASCIIArt
func TestHandleASCIIArt(t *testing.T) {
	// Create a request to pass to our handler
	form := url.Values{}
	form.Set("banner", "standard.txt")
	form.Set("input", "Hello")

	req, err := http.NewRequest("POST", "/ascii-art", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleASCIIArt)

	// Call the handler directly with our request and response recorder
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	file, err := os.ReadFile("testCases/expectedOutput1.txt")
	if err != nil {
		t.Fatalf("Error reading %s", "expectedOutput1.txt")
	}

	// Check the response body to ensure the expected ASCII art result is generated
	expected := string(file)
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Test if MethodNotAllowed is handled correctly
func TestHandleASCIIArt_MethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("GET", "/ascii-art", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleASCIIArt)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

// Test if the case of missing inputs is handled correctly
func TestHandleASCIIArt_MissingInputs(t *testing.T) {
	req, err := http.NewRequest("POST", "/ascii-art", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleASCIIArt)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// Test if illegal characters in the input are handled correctly
func TestHandleASCIIArt_IllegalCharacters(t *testing.T) {
	form := url.Values{}
	form.Set("banner", "standard.txt")
	form.Set("input", "Hello world \u20AC") // Input with illegal character

	req, err := http.NewRequest("POST", "/ascii-art", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleASCIIArt)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
