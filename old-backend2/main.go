package main

import (
	"context"
	"delta/handlers"
	"delta/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	if err := handlers.InitWebAuthn(); err != nil {
		log.Fatalf("Failed to initialize WebAuthn: %v", err)
	}

	handlers.InitSessionStore()

	// Initialize database
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	marketDataUrl := os.Getenv("MARKET_DATA_URL")
	routingUrl := os.Getenv("ROUTING_URL")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the models
	db.AutoMigrate(&models.Dashboard{}, &models.Panel{})

	routingClient := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   routingUrl,
	})

	marketDataClient := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   marketDataUrl,
	})

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

	// Dashboard routes
	dashboard := api.PathPrefix("/dashboard").Subrouter()
	dashboard.HandleFunc("", dashboardHandler.Create).Methods("POST")
	dashboard.HandleFunc("", dashboardHandler.GetAll).Methods("GET")
	dashboard.HandleFunc("/{id}", dashboardHandler.GetByID).Methods("GET")
	dashboard.HandleFunc("/{id}", dashboardHandler.Update).Methods("PUT")
	dashboard.HandleFunc("/{id}", dashboardHandler.Delete).Methods("DELETE")
	dashboard.HandleFunc("/{id}/panel", dashboardHandler.UpdatePanel).Methods("PUT")

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
