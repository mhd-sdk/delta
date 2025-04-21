package models

import (
	"crypto/rand"
	"encoding/binary"

	"github.com/duo-labs/webauthn/webauthn"
)

// User represents the user model
type User struct {
	ID          uint64
	Username    string
	DisplayName string
	Email       string
	credentials []webauthn.Credential
}

// NewUser creates and returns a new User
func NewUser(name string, displayName string) *User {

	user := &User{}
	user.ID = randomUint64()
	user.Username = name
	user.DisplayName = displayName
	// user.credentials = []webauthn.Credential{}

	return user
}

func randomUint64() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}

// WebAuthnID returns the user's ID
func (u User) WebAuthnID() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, uint64(u.ID))
	return buf
}

// WebAuthnName returns the user's username
func (u User) WebAuthnName() string {
	return u.Username
}

// WebAuthnDisplayName returns the user's display name
func (u User) WebAuthnDisplayName() string {
	return u.DisplayName
}

// WebAuthnIcon is not (yet) implemented
func (u User) WebAuthnIcon() string {
	return ""
}

// AddCredential associates the credential to the user
func (u *User) AddCredential(cred webauthn.Credential) {
	u.credentials = append(u.credentials, cred)
}

// WebAuthnCredentials returns credentials owned by the user
func (u User) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}
