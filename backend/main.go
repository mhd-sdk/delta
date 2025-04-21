package main

import (
	"context"
	"delta/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	if err := handlers.InitWebAuthn(); err != nil {
		log.Fatalf("Failed to initialize WebAuthn: %v", err)
	}

	handlers.InitSessionStore()

	// Create a new router
	r := mux.NewRouter()

	// Apply global middlewares
	r.Use(handlers.LoggingMiddleware)
	r.Use(handlers.RecoveryMiddleware)
	r.Use(handlers.CorsMiddleware)

	// API Routes
	api := r.PathPrefix("/api").Subrouter()

	// Auth routes
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register/begin", handlers.BeginRegistration).Methods("POST")
	auth.HandleFunc("/register/finish", handlers.FinishRegistration).Methods("POST")
	auth.HandleFunc("/login/begin", handlers.BeginLogin).Methods("POST")
	auth.HandleFunc("/login/finish", handlers.FinishLogin).Methods("POST")
	auth.HandleFunc("/logout", handlers.Logout).Methods("POST")
	
	// Don't forget to add OPTIONS for preflight requests
	auth.HandleFunc("/register/begin", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	auth.HandleFunc("/register/finish", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	auth.HandleFunc("/login/begin", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	auth.HandleFunc("/login/finish", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	auth.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	// Health check route
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Create server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start the server in a goroutine
	go func() {
		log.Printf("Server started on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
