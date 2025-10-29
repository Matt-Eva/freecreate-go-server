package pgModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"not_null;uniqueIndex"`
	Birthday time.Time `gorm:"not_null"`
	SessionUUID uuid.UUID `gorm:"index"`
	Username string
}
