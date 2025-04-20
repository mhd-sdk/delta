package models

import (
	"context"
	"time"

	"delta/backend/db"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// WebAuthn fields
	Credentials []Credential `json:"-"`
}

type Credential struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	CredentialID    string    `json:"credential_id"`
	PublicKey       []byte    `json:"public_key"`
	AttestationType string    `json:"attestation_type"`
	AAGUID          string    `json:"aaguid"`
	SignCount       uint32    `json:"sign_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// WebAuthn implementation methods

func (u *User) WebAuthnID() []byte {
	return u.ID[:]
}

func (u *User) WebAuthnName() string {
	return u.Username
}

func (u *User) WebAuthnDisplayName() string {
	return u.Username
}

func (u *User) WebAuthnIcon() string {
	return ""
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	credentials := make([]webauthn.Credential, len(u.Credentials))
	for i, cred := range u.Credentials {
		credentials[i] = webauthn.Credential{
			ID:              []byte(cred.CredentialID),
			PublicKey:       cred.PublicKey,
			AttestationType: cred.AttestationType,
			Authenticator: webauthn.Authenticator{
				AAGUID:    []byte(cred.AAGUID),
				SignCount: cred.SignCount,
			},
		}
	}
	return credentials
}

func GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	user := &User{ID: id}
	err := db.DB.QueryRow(ctx, `
		SELECT username, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`, id).Scan(&user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	user.Credentials, err = getUserCredentials(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	user := &User{}
	err := db.DB.QueryRow(ctx, `
		SELECT id, username, email, created_at, updated_at
		FROM users
		WHERE username = $1
	`, username).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	user.Credentials, err = getUserCredentials(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(ctx context.Context, username, email string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
	}

	err := db.DB.QueryRow(ctx, `
		INSERT INTO users (username, email)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`, username, email).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) AddCredential(ctx context.Context, c *webauthn.Credential) error {
	credentialID := string(c.ID)

	_, err := db.DB.Exec(ctx, `
		INSERT INTO webauthn_credentials 
		(user_id, credential_id, public_key, attestation_type, aaguid, sign_count)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, u.ID, credentialID, c.PublicKey, c.AttestationType, c.Authenticator.AAGUID, c.Authenticator.SignCount)

	return err
}

func (u *User) UpdateCredential(ctx context.Context, c *webauthn.Credential) error {
	credentialID := string(c.ID)

	_, err := db.DB.Exec(ctx, `
		UPDATE webauthn_credentials
		SET sign_count = $1, updated_at = NOW()
		WHERE user_id = $2 AND credential_id = $3
	`, c.Authenticator.SignCount, u.ID, credentialID)

	return err
}

func getUserCredentials(ctx context.Context, userID uuid.UUID) ([]Credential, error) {
	rows, err := db.DB.Query(ctx, `
		SELECT id, user_id, credential_id, public_key, attestation_type, aaguid, sign_count, created_at, updated_at
		FROM webauthn_credentials
		WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []Credential
	for rows.Next() {
		var cred Credential
		if err := rows.Scan(
			&cred.ID, &cred.UserID, &cred.CredentialID, &cred.PublicKey,
			&cred.AttestationType, &cred.AAGUID, &cred.SignCount,
			&cred.CreatedAt, &cred.UpdatedAt,
		); err != nil {
			return nil, err
		}
		credentials = append(credentials, cred)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return credentials, nil
}
