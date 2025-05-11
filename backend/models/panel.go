package models

import (
	"time"

	"github.com/google/uuid"
)

type Panel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	Position  int       `json:"position"`
	Size      string    `json:"size"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Dashboard struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Panels    []Panel   `json:"panels" gorm:"many2many:dashboard_panels;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
