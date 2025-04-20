package handlers

import (
	"delta/backend/models"
	"delta/backend/services"
	"errors"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}

func BeginRegistration(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return sendErrorResponse(c, "invalid request", fiber.StatusBadRequest)
	}

	if req.Username == "" || req.Email == "" {
		return sendErrorResponse(c, "username and email are required", fiber.StatusBadRequest)
	}

	// Check if user already exists
	existingUser, err := models.GetUserByUsername(c.Context(), req.Username)
	if err == nil && existingUser != nil {
		// User exists, begin authentication instead
		return beginAuthentication(c, existingUser)
	} else if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return sendErrorResponse(c, "database error", fiber.StatusInternalServerError)
	}

	// Create new user
	user, err := models.CreateUser(c.Context(), req.Username, req.Email)
	if err != nil {
		return sendErrorResponse(c, "failed to create user", fiber.StatusInternalServerError)
	}

	// Generate WebAuthn credential creation options
	options, sessionData, err := services.WebAuthn.BeginRegistration(user)
	if err != nil {
		return sendErrorResponse(c, "failed to begin registration: "+err.Error(), fiber.StatusInternalServerError)
	}

	// Store registration session data
	sessionID := uuid.New().String()
	services.Sessions[sessionID] = services.SessionData{
		Registration: sessionData,
	}

	// Set session cookie
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "lax",
		MaxAge:   int(time.Hour.Seconds()),
	})

	return c.JSON(options)
}

func FinishRegistration(c *fiber.Ctx) error {
	// Get session cookie
	sessionID := c.Cookies("session_id")
	if sessionID == "" {
		return sendErrorResponse(c, "session not found", fiber.StatusBadRequest)
	}

	sessionData, exists := services.Sessions[sessionID]
	if !exists || sessionData.Registration == nil {
		return sendErrorResponse(c, "registration session not found", fiber.StatusBadRequest)
	}

	// Get user from the session data
	user, err := models.GetUserByID(c.Context(), uuid.UUID(sessionData.Registration.UserID))
	if err != nil {
		return sendErrorResponse(c, "user not found", fiber.StatusInternalServerError)
	}

	// Finish registration using our helper function
	credential, err := services.FinishWebAuthnRegistration(user, *sessionData.Registration, c.Body())
	if err != nil {
		return sendErrorResponse(c, "failed to complete registration: "+err.Error(), fiber.StatusBadRequest)
	}

	// Store credential in database
	if err := user.AddCredential(c.Context(), credential); err != nil {
		return sendErrorResponse(c, "failed to store credential", fiber.StatusInternalServerError)
	}

	// Delete session
	delete(services.Sessions, sessionID)

	// Generate JWT
	token, err := services.GenerateJWT(user)
	if err != nil {
		return sendErrorResponse(c, "failed to generate token", fiber.StatusInternalServerError)
	}

	expiryTime, _ := services.GetTokenExpiryTime(token)

	// Return the token
	return c.JSON(AuthResponse{
		Token:     token,
		ExpiresAt: expiryTime,
		Username:  user.Username,
		Email:     user.Email,
	})
}

func BeginLogin(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return sendErrorResponse(c, "invalid request", fiber.StatusBadRequest)
	}

	if req.Username == "" {
		return sendErrorResponse(c, "username is required", fiber.StatusBadRequest)
	}

	user, err := models.GetUserByUsername(c.Context(), req.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sendErrorResponse(c, "user not found", fiber.StatusNotFound)
		} else {
			return sendErrorResponse(c, "database error", fiber.StatusInternalServerError)
		}
	}

	return beginAuthentication(c, user)
}

func beginAuthentication(c *fiber.Ctx, user *models.User) error {
	options, sessionData, err := services.WebAuthn.BeginLogin(user)
	if err != nil {
		return sendErrorResponse(c, "failed to begin authentication: "+err.Error(), fiber.StatusInternalServerError)
	}

	// Store authentication session data
	sessionID := uuid.New().String()
	services.Sessions[sessionID] = services.SessionData{
		Authentication: sessionData,
	}

	// Set session cookie
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "lax",
		MaxAge:   int(time.Hour.Seconds()),
	})

	return c.JSON(options)
}

func FinishLogin(c *fiber.Ctx) error {
	// Get session cookie
	sessionID := c.Cookies("session_id")
	if sessionID == "" {
		return sendErrorResponse(c, "session not found", fiber.StatusBadRequest)
	}

	sessionData, exists := services.Sessions[sessionID]
	if !exists || sessionData.Authentication == nil {
		return sendErrorResponse(c, "authentication session not found", fiber.StatusBadRequest)
	}

	// Get user from session data
	user, err := models.GetUserByID(c.Context(), uuid.UUID(sessionData.Authentication.UserID))
	if err != nil {
		return sendErrorResponse(c, "user not found", fiber.StatusInternalServerError)
	}

	// Finish login using our helper function
	credential, err := services.FinishWebAuthnLogin(user, *sessionData.Authentication, c.Body())
	if err != nil {
		return sendErrorResponse(c, "failed to complete authentication: "+err.Error(), fiber.StatusBadRequest)
	}

	// Update credential sign count in database
	if err := user.UpdateCredential(c.Context(), credential); err != nil {
		return sendErrorResponse(c, "failed to update credential", fiber.StatusInternalServerError)
	}

	// Delete session
	delete(services.Sessions, sessionID)

	// Generate JWT
	token, err := services.GenerateJWT(user)
	if err != nil {
		return sendErrorResponse(c, "failed to generate token", fiber.StatusInternalServerError)
	}

	expiryTime, _ := services.GetTokenExpiryTime(token)

	// Return the token
	return c.JSON(AuthResponse{
		Token:     token,
		ExpiresAt: expiryTime,
		Username:  user.Username,
		Email:     user.Email,
	})
}

func VerifyToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return sendErrorResponse(c, "authorization header required", fiber.StatusUnauthorized)
	}

	// Extract the token
	tokenString := authHeader[7:] // Remove "Bearer " prefix

	// Verify the token
	user, err := services.GetUserFromToken(c.Context(), tokenString)
	if err != nil {
		return sendErrorResponse(c, "invalid token", fiber.StatusUnauthorized)
	}

	// Check if token is expired
	if services.IsTokenExpired(tokenString) {
		return sendErrorResponse(c, "token expired", fiber.StatusUnauthorized)
	}

	expiryTime, _ := services.GetTokenExpiryTime(tokenString)

	// Return user information
	return c.JSON(AuthResponse{
		Token:     tokenString,
		ExpiresAt: expiryTime,
		Username:  user.Username,
		Email:     user.Email,
	})
}

// Utility function to send error responses
func sendErrorResponse(c *fiber.Ctx, message string, statusCode int) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Error: message,
	})
}

// A fake implementation to bridge between the parsed response and what webauthn expects
type fakeResponse struct {
	parsedCredCreation  *protocol.ParsedCredentialCreationData
	parsedCredAssertion *protocol.ParsedCredentialAssertionData
}

// Implement the necessary methods to satisfy the interface
func (f *fakeResponse) ParsedAttestationResponse() *protocol.ParsedCredentialCreationData {
	return f.parsedCredCreation
}

func (f *fakeResponse) ParsedAssertionResponse() *protocol.ParsedCredentialAssertionData {
	return f.parsedCredAssertion
}
