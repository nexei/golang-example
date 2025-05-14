//go:build e2e
// +build e2e

package main

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestE2E(t *testing.T) {
	// Wait for server to start
	time.Sleep(2 * time.Second)

	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	expected := "Hello, World!"
	if response.Message != expected {
		t.Errorf("Expected message %q; got %q", expected, response.Message)
	}
}

func TestUserProfileE2E(t *testing.T) {
	// Wait for server to start
	time.Sleep(2 * time.Second)

	resp, err := http.Get("http://localhost:8080/user/profile")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
	}

	var profile UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		t.Fatal(err)
	}

	// This assertion will fail
	if profile.Email != "different@example.com" {
		t.Errorf("Expected email to be 'different@example.com', got %s", profile.Email)
	}
} 