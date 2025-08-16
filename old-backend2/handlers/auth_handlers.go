package handlers

import (
	"delta/models"
	"delta/repositories"
	"delta/sessionstore"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	WebAuthn *webauthn.WebAuthn
)

func InitWebAuthn() error {
	displayName := os.Getenv("WEBAUTHN_DISPLAY_NAME")
	domain := os.Getenv("WEBAUTHN_DOMAIN")
	origin := os.Getenv("WEBAUTHN_ORIGIN")

	wconfig := &webauthn.Config{
		RPDisplayName: displayName,
		RPID:          domain,
		RPOrigins:     []string{origin},
	}

	var err error
	WebAuthn, err = webauthn.New(wconfig)
	if err != nil {
		return fmt.Errorf("failed to create WebAuthn: %w", err)
	}

	return nil
}

var SessionStore *sessionstore.SessionStore

func InitSessionStore() {
	SessionStore = sessionstore.NewSessionStore()
}

func BeginRegistration(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request
	var requestData struct {
		Username string `json:"username"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		jsonResponse(w, map[string]string{"error": "Invalid request format"}, http.StatusBadRequest)
		return
	}

	if requestData.Username == "" {
		jsonResponse(w, map[string]string{"error": "Username is required"}, http.StatusBadRequest)
		return
	}

	// Get user
	user, err := repositories.DB().GetUserByName(requestData.Username)
	// User doesn't exist, create new user
	if err != nil {
		displayName := strings.Split(requestData.Username, "@")[0]
		user = models.NewUser(requestData.Username, displayName)
		repositories.DB().PutUser(user)
	} else {
		jsonResponse(w, map[string]string{"error": "User already exists"}, http.StatusBadRequest)
		return
	}

	// Generate PublicKeyCredentialCreationOptions, session data
	options, sessionData, err := WebAuthn.BeginRegistration(user)
	if err != nil {
		log.Println(err)
		jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	// Store session data as marshaled JSON
	SessionStore.Set("registration", *sessionData)
	if err != nil {
		log.Println(err)
		jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, options, http.StatusOK)
}

func FinishRegistration(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	// Get the user
	user, err := repositories.DB().GetUserByName(username)
	if err != nil {
		jsonResponse(w, map[string]string{"error": "User not found"}, http.StatusBadRequest)
		return
	}

	// load the session data
	sessionData, ok := SessionStore.Get("registration")
	if ok != true {
		log.Println(err)
		jsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	credential, err := WebAuthn.FinishRegistration(user, sessionData.(webauthn.SessionData), r)
	if err != nil {
		log.Println(err)
		jsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.AddCredential(*credential)

	response := map[string]interface{}{
		"status":  "Registration successful",
		"user":    user,
		"session": sessionData,
	}

	jsonResponse(w, response, http.StatusOK)

}

func BeginLogin(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request
	var requestData struct {
		Username string `json:"username"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		jsonResponse(w, map[string]string{"error": "Invalid request format"}, http.StatusBadRequest)
		return
	}

	if requestData.Username == "" {
		jsonResponse(w, map[string]string{"error": "Username is required"}, http.StatusBadRequest)
		return
	}

	// Get the user
	user, err := repositories.DB().GetUserByName(requestData.Username)
	if err != nil {
		jsonResponse(w, map[string]string{"error": "User not found"}, http.StatusBadRequest)
		return
	}

	// Begin WebAuthn login
	options, sessionData, err := WebAuthn.BeginLogin(user)
	if err != nil {
		jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	// Store session data
	SessionStore.Set("authentication", *sessionData)

	jsonResponse(w, options, http.StatusOK)
}

func FinishLogin(w http.ResponseWriter, r *http.Request) {
	// get username from query params
	username := r.URL.Query().Get("username")
	user, err := repositories.DB().GetUserByName(username)
	if err != nil {
		jsonResponse(w, map[string]string{"error": "User not found"}, http.StatusBadRequest)
		return
	}

	// Get the session data
	sessionData, ok := SessionStore.Get("authentication")
	if ok != true {
		jsonResponse(w, map[string]string{"error": "Session not found"}, http.StatusBadRequest)
		return
	}

	credential, err := WebAuthn.FinishLogin(user, sessionData.(webauthn.SessionData), r)
	if err != nil {
		jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	// Generate authentication response
	authResponse := map[string]interface{}{
		"status":     "Login successful",
		"username":   user.Username,
		"credential": credential,
	}

	jsonResponse(w, authResponse, http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}
