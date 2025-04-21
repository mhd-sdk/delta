package repositories

import (
	"delta/models"
	"fmt"
	"sync"
)

type userdb struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

var db *userdb

// DB returns a userdb singleton
func DB() *userdb {
	if db == nil {
		db = &userdb{
			users: make(map[string]*models.User),
		}
	}
	return db
}

// GetUserByName returns a *User by the user's username
func (db *userdb) GetUserByName(name string) (*models.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	user, ok := db.users[name]
	if !ok {
		return nil, fmt.Errorf("error getting user '%s': does not exist", name)
	}
	return user, nil
}

// GetUserByID returns a *User by the user's ID
func (db *userdb) GetUserByID(id uint64) (*models.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, user := range db.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("error getting user with ID '%s': does not exist", id)
}

// PutUser stores a new user by the user's username
func (db *userdb) PutUser(user *models.User) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.Username] = user
}

// UpdateUser updates an existing user
func (db *userdb) UpdateUser(user *models.User) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, ok := db.users[user.Username]; !ok {
		return fmt.Errorf("error updating user '%s': does not exist", user.Username)
	}
	db.users[user.Username] = user
	return nil
}
