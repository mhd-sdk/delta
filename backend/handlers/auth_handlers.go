package handlers

import (
	"delta/models"
	"delta/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/duo-labs/webauthn.io/session"
	"github.com/duo-labs/webauthn/webauthn"
)

var (
	WebAuthn *webauthn.WebAuthn
)

func InitWebAuthn() error {
	displayName := os.Getenv("WEBAUTHN_DISPLAY_NAME")
	domain := os.Getenv("WEBAUTHN_DOMAIN")

	wconfig := &webauthn.Config{
		RPDisplayName: displayName,
		RPID:          domain,
	}

	var err error
	WebAuthn, err = webauthn.New(wconfig)
	if err != nil {
		return fmt.Errorf("failed to create WebAuthn: %w", err)
	}

	return nil
}

var SessionStore *session.Store

func InitSessionStore() {
	var err error
	SessionStore, err = session.NewStore()
	if err != nil {
		log.Fatal("failed to create session store:", err)
	}
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
	}

	// Generate PublicKeyCredentialCreationOptions, session data
	options, sessionData, err := WebAuthn.BeginRegistration(
		user,
	)

	if err != nil {
		log.Println(err)
		jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	// Store session data as marshaled JSON
	err = SessionStore.SaveWebauthnSession("registration", sessionData, r, w)
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

	//  continuer ici...

	jsonResponse(w, map[string]string{"status": "Registration successful"}, http.StatusOK)
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
	err = SessionStore.SaveWebauthnSession("authentication", sessionData, r, w)
	if err != nil {
		jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, options, http.StatusOK)
}

func FinishLogin(w http.ResponseWriter, r *http.Request) {
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

	// Get the session data
	sessionData, err := SessionStore.GetWebauthnSession("authentication", r)
	if err != nil {
		jsonResponse(w, map[string]string{"error": "Session not found"}, http.StatusBadRequest)
		return
	}

	// Complete the login
	_, err = WebAuthn.FinishLogin(user, sessionData, r)
	if err != nil {
		jsonResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	// Create session for the user
	session, err := SessionStore.New(r, "auth-session")
	if err != nil {
		jsonResponse(w, map[string]string{"error": "Failed to create session"}, http.StatusInternalServerError)
		return
	}

	// Save the username in the session
	session.Values["username"] = user.Username
	if err := session.Save(r, w); err != nil {
		jsonResponse(w, map[string]string{"error": "Failed to save session"}, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]string{"status": "Login successful"}, http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Get the session
	session, err := SessionStore.Get(r, "auth-session")
	if err != nil {
		jsonResponse(w, map[string]string{"error": "Not logged in"}, http.StatusBadRequest)
		return
	}

	// Clear the session
	session.Values = map[interface{}]interface{}{}
	if err := session.Save(r, w); err != nil {
		jsonResponse(w, map[string]string{"error": "Failed to logout"}, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]string{"status": "Logout successful"}, http.StatusOK)
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
