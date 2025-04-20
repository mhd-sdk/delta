package services

import (
	"bytes"
	"context"
	"delta/backend/models"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	WebAuthn *webauthn.WebAuthn
)

type SessionData struct {
	Registration   *webauthn.SessionData
	Authentication *webauthn.SessionData
}

// Sessions store in memory for now, in production use Redis/etc.
var Sessions = make(map[string]SessionData)

func InitWebAuthn() error {
	displayName := os.Getenv("WEBAUTHN_DISPLAY_NAME")
	domain := os.Getenv("WEBAUTHN_DOMAIN")
	origin := os.Getenv("WEBAUTHN_ORIGIN")

	wconfig := &webauthn.Config{
		RPDisplayName: displayName,
		RPID:          domain,
		RPOrigin:      origin,
	}

	var err error
	WebAuthn, err = webauthn.New(wconfig)
	if err != nil {
		return fmt.Errorf("failed to create WebAuthn: %w", err)
	}

	return nil
}

// Wrap FinishRegistration to handle the Fiber body bytes instead of http.Request
func FinishWebAuthnRegistration(user webauthn.User, session webauthn.SessionData, body []byte) (*webauthn.Credential, error) {
	bodyReader := bytes.NewReader(body)
	parsedResponse, err := protocol.ParseCredentialCreationResponseBody(bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse credential: %w", err)
	}

	return WebAuthn.CreateCredential(user, session, parsedResponse)
}

// Wrap FinishLogin to handle the Fiber body bytes instead of http.Request
func FinishWebAuthnLogin(user webauthn.User, session webauthn.SessionData, body []byte) (*webauthn.Credential, error) {
	bodyReader := bytes.NewReader(body)
	parsedResponse, err := protocol.ParseCredentialRequestResponseBody(bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse credential: %w", err)
	}

	return WebAuthn.ValidateLogin(user, session, parsedResponse)
}

func GenerateJWT(user *models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	sessionDurationStr := os.Getenv("SESSION_DURATION_HOURS")
	sessionDuration, err := strconv.Atoi(sessionDurationStr)
	if err != nil {
		sessionDuration = 2 // Default to 2 hours
	}

	expirationTime := time.Now().Add(time.Duration(sessionDuration) * time.Hour)

	claims := jwt.MapClaims{
		"sub":      user.ID.String(),
		"username": user.Username,
		"email":    user.Email,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func ParseJWT(tokenString string) (*jwt.MapClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GetUserFromToken(ctx context.Context, tokenString string) (*models.User, error) {
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	userIDStr, ok := (*claims)["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid user ID in token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	user, err := models.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func IsTokenExpired(tokenString string) bool {
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return true
	}

	exp, ok := (*claims)["exp"].(float64)
	if !ok {
		return true
	}

	return time.Now().Unix() > int64(exp)
}

func GetTokenExpiryTime(tokenString string) (time.Time, error) {
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return time.Time{}, err
	}

	exp, ok := (*claims)["exp"].(float64)
	if !ok {
		return time.Time{}, fmt.Errorf("invalid expiration time in token")
	}

	return time.Unix(int64(exp), 0), nil
}
