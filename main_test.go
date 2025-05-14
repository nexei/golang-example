package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response Response
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	expected := "Hello, World!"
	if response.Message != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Message, expected)
	}
}

func TestUserProfileHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userProfileHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var profile UserProfile
	if err := json.NewDecoder(rr.Body).Decode(&profile); err != nil {
		t.Fatal(err)
	}

	// This assertion will fail
	if profile.ID != 2 {
		t.Errorf("Expected user ID to be 2, got %d", profile.ID)
	}

	// This assertion will pass
	if profile.Username != "testuser" {
		t.Errorf("Expected username to be 'testuser', got %s", profile.Username)
	}
} 