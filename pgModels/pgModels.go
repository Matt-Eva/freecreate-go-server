package pgModels

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string    `gorm:"not_null;uniqueIndex"`
	Birthday  time.Time `gorm:"not_null"`
	// we reset this uuid each time a user logs in to help protect user security from
	// session to session.
	SessionUUID uuid.UUID `gorm:"index"`
	Username    string
}
