package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func Test_set_and_get_api(t *testing.T) {
	// Set a user by using http POST method to /users
	// Get the user by using http GET method to /users/{id}
	// Check if the user is the same as the one set

	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// remove all keys from the current database
	redisClient.FlushAll(context.Background())
	restServer := initRestServer("8080", redisClient)
	// Test POST /users endpoint
	user := UserInput{Name: "John Doe", Age: 10}
	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	rr := httptest.NewRecorder()
	restServer.ServeHTTP(rr, req)

	// Check if the POST request was successful
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Get the user by using http GET method to /users/{id}
	// Extract the user ID from the response body
	var response struct {
		ID string `json:"id"`
	}
	json.Unmarshal(rr.Body.Bytes(), &response)

	// Test GET /users/{id} endpoint
	req, _ = http.NewRequest("GET", fmt.Sprintf("/users/%s", response.ID), nil)
	rr = httptest.NewRecorder()
	restServer.ServeHTTP(rr, req)

	// Check if the GET request was successful
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check if the retrieved user is the same as the one set
	var retrievedUser UserOutput
	// log the response body as tring
	log.Debug().Msgf("Response body: %s", rr.Body.String())
	// log response body as raw json string
	log.Debug().Msgf("Response body: %s", string(rr.Body.Bytes()))
	jsonData := string(rr.Body.Bytes())
	fmt.Print(jsonData)
	json.Unmarshal(rr.Body.Bytes(), &retrievedUser)
	log.Debug().Msgf("Retrieved user: %+v", retrievedUser)
	if retrievedUser.Name != user.Name || retrievedUser.Age != user.Age {
		t.Errorf("Expected user %+v, but got %+v", user, retrievedUser)
	}

}

func Test_debug_jsin(t *testing.T) {
	// Sample JSON data
	jsonData := `{"age":10,"id":"de06118e-7286-4843-948a-4a74eb8fb789","name":"John Doe"}`

	// Unmarshal the JSON data into the UserOutput struct
	var user UserOutput
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the struct
	fmt.Printf("UserOutput struct: %+v\n", user)
}
