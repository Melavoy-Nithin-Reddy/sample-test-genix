package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sample-user-service/handlers"
	"github.com/sample-user-service/repository"
	"github.com/sample-user-service/service"
)

func main() {
	// Initialize repository
	userRepo := repository.NewUserRepository()

	// Initialize service
	userService := service.NewUserService(userRepo)

	// Initialize handler
	userHandler := handlers.NewUserHandler(userService)

	// Setup router
	router := mux.NewRouter()

	// Health check endpoint
	router.HandleFunc("/health", userHandler.HealthCheck).Methods("GET")

	// User endpoints
	router.HandleFunc("/api/v1/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/users/stats", userHandler.GetUserStats).Methods("GET")

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting User Service on %s", addr)
	log.Printf("Health check available at http://localhost%s/health", addr)
	log.Printf("API endpoints available at http://localhost%s/api/v1/users", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// Made with Bob
