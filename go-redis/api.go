package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"math/rand/v2"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type RestServer struct {
	router      *chi.Mux
	redisClient *redis.Client
}

type UserInput struct {
	Name string `json:"name" redis:"name"`
	Age  int    `json:"age" redis:"age"`
}

type UserOutput struct {
	ID   string `json:"id" redis:"id"`
	Name string `json:"name" redis:"name"`
	Age  int    `json:"age" redis:"age"`
}

func (s *RestServer) createUser(w http.ResponseWriter, r *http.Request) {
	// Handle create user logic here
	// convert request body to UserInput struct
	var input UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// generate random unique ID
	uid := uuid.New().String()

	// convert UserInput to UserOutput
	user := UserOutput{
		ID:   uid,
		Name: input.Name,
		Age:  input.Age,
	}

	// Convert UserOutput to map[string]interface{} and write to Redis
	// use HSet to store user data in Redis

	err = s.redisClient.HSet(context.Background(), uid, user).Err()
	if err != nil {
		// use zerolog to log error
		log.Error().Err(err).Msg("failed to write user to redis")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return UserOutput as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *RestServer) getUser(w http.ResponseWriter, r *http.Request) {
	// Handle get user logic here
	// extract user ID from URL path
	uid := chi.URLParam(r, "id")
	var user UserOutput

	// get from Redis using HGetAll and return as JSON response UserOutput
	err := s.redisClient.HGetAll(context.Background(), uid).Scan(&user)

	if err != nil {
		log.Error().Err(err).Msg("failed to get user from redis")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	log.Debug().Msg(fmt.Sprintf("User: %v", user))
	json.NewEncoder(w).Encode(user)
}

func (s *RestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *RestServer) randomFail(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// Recover from panic and return internal server error
		if r := recover(); r != nil {
			log.Error().Msgf("Recovered from panic: %v", r)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}()
	// Randomly panic to simulate server failure
	if rand.IntN(10) > 5 {
		panic("random failure")
	}
	w.Write([]byte("success"))
}

func (s *RestServer) alwaysFail(w http.ResponseWriter, r *http.Request) {
	// always fail endpoint to check Exception handling
	panic("alwaysFail function failure")
}

// define recover middleware to catch panic and return 500
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Recover from panic and return internal server error
			if r := recover(); r != nil {
				log.Error().Msgf("Recovered from panic: %v", r)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func initRestServer(redisClient *redis.Client) *RestServer {
	chiRouter := chi.NewRouter()
	chiRouter.Use(recoverMiddleware)

	server := &RestServer{
		router:      chiRouter,
		redisClient: redisClient,
	}
	server.router.Post("/users", server.createUser)
	server.router.Get("/users/{id}", server.getUser)
	server.router.Get("/random-fail", server.randomFail)
	server.router.Get("/always-fail", server.alwaysFail)

	return server
}

func main() {
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	server := initRestServer(redisClient)
	log.Info().Msg("Starting server on :8080")
	http.ListenAndServe(":8080", server.router)
	log.Info().Msg("Server stopped")
}
