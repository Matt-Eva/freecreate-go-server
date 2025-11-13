package pgModels

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string    `gorm:"not_null;uniqueIndex"`
	Birthday  time.Time `gorm:"not_null"`
	// we reset this uuid each time a user logs in to help protect user security from
	// session to session.
	SessionUUID uuid.UUID `gorm:"index:idx_user_session_uuid"`
	Username    string
}

type Creator struct {
	ID uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID uuid.UUID `gorm:"index:idx_creator_uuid"`
	Name string `gorm:"not_null"`
	UserID uint
	User User
}

type Writing struct {
	ID uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	WritingType string
	UUID uuid.UUID `gorm:"index:idx_writing_uuid"`
	Title string
	Tags pq.StringArray `gorm:"type:text[];index:idx_writing_tags"`
	UserID uint
	User User
	CreatorID uint
	Creator Creator
}

type Content struct {
	ID uint `gorm:"primaryKey"`
	UUID uuid.UUID `gorm:"index:idx_content_uuid"`
	Title string
	WritingID uint
	Writing Writing
}