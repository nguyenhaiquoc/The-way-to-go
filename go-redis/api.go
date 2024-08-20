package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"go-redis/internal/models"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"

	"math/rand/v2"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type RestServer struct {
	httpServer  *http.Server
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
		zlog.Error().Err(err).Msg("failed to write user to redis")
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
		zlog.Error().Err(err).Msg("failed to get user from redis")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	zlog.Debug().Msg(fmt.Sprintf("User: %v", user))
	json.NewEncoder(w).Encode(user)
}

func (s *RestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.httpServer.Handler.ServeHTTP(w, r)
}

func (s *RestServer) randomFail(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// Recover from panic and return internal server error
		if r := recover(); r != nil {
			zlog.Error().Msgf("Recovered from panic: %v", r)
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

func recoverMiddleware(next http.Handler) http.Handler {
	/*
		RecoverMiddleware is a middleware that recovers from panic and returns internal server error
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Recover from panic and return internal server error
			if r := recover(); r != nil {
				zlog.Error().Msgf("Recovered from panic: %v", r)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (s *RestServer) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Post("/users", s.createUser)
	r.Get("/users/{id}", s.getUser)
	r.Get("/random-fail", s.randomFail)
	r.Get("/always-fail", s.alwaysFail)
	return r
}

type zerologWriter struct {
	logger zerolog.Logger
}

func (w zerologWriter) Write(p []byte) (n int, err error) {
	w.logger.Error().Msg(string(p))
	return len(p), nil
}

func initRestServer(addr string, redisClient *redis.Client, userModel *models.UserModel) *RestServer {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	zWriter := zerologWriter{logger: logger}

	server := &RestServer{
		httpServer:  &http.Server{Addr: ":" + addr, ErrorLog: log.New(zWriter, "", 0)},
		redisClient: redisClient,
	}
	server.httpServer.Handler = server.routes()
	return server
}

func getPostgreConnection(dsn string) (*pgx.Conn, error) {
	// get the postgre connection
	// dsn := "postgres://your_user:your_password@localhost:5432/your_db?sslmode=disable"
	// use pgx driver to connect to postgre (no )
	zlog.Info().Msg("Connecting to Postgres")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		zlog.Error().Err(err).Msg("Failed to ping Postgres")
		return nil, err
	}
	return conn, nil
}

func main() {
	dsn := flag.String("dsn", "postgres://your_user:your_password@localhost:5432/your_db?sslmode=disable", "Postgres DSN")
	dbConn, err := getPostgreConnection(*dsn)
	if err != nil {
		zlog.Error().Err(err).Msg("Failed to connect to Postgres")
		// Exit with error
		os.Exit(1)
	}
	defer dbConn.Close(context.Background())

	userModel := &models.UserModel{DB: dbConn}

	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	server := initRestServer("8080", redisClient, userModel)
	zlog.Info().Msg("Starting server on :8080")

	// declare postgres dsn from flag

	err = server.httpServer.ListenAndServe()
	if err != nil {
		zlog.Error().Err(err).Msg("Failed to start server")
	}
	zlog.Info().Msg("Server stopped")
}
