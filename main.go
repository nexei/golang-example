package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type UserProfile struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Simulate a user profile
	profile := UserProfile{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user/profile", userProfileHandler)
	
	log.Printf("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
} 