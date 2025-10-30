package pgModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `gorm:"not_null;uniqueIndex"`
	Birthday time.Time `gorm:"not_null"`
	// we reset this uuid each time a user logs in to help protect user security from
	// session to session.
	SessionUUID uuid.UUID `gorm:"index"`
	Username    string
}
