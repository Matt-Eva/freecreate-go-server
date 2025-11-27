package pgModels

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Email       string    `gorm:"not_null;uniqueIndex"`
	Birthday    time.Time `gorm:"not_null"`
	SessionUUID uuid.UUID `gorm:"index:idx_user_session_uuid"` // we reset this uuid each time a user logs in to help protect user security from session to session.
}

type Bookshelf struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	UserID    uint
	User      User
}

type Creator struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID      uuid.UUID `gorm:"index:idx_creator_uuid"`
	Name      string    `gorm:"not_null"`
	UserID    uint
	User      User
	Tags pq.StringArray `gorm:"type:text[];index:idx_creator_tags"`
	Followers uint
	Subscribers uint
	RecurringDonations uint
	Donations uint
	Views uint
	Likes uint
}

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not_null;uniqueIndex"`
}

type ContentTag struct {
	TagID     uint `gorm:"primaryKey"`
	Tag       Tag
	WritingID uint `gorm:"primaryKey"`
	Writing   Writing
	CreatorID uint `gorm:"primaryKey"`
	Creator Creator
	CreatedAt time.Time
}

type Writing struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	WritingType    string    `gorm:"not_null"`
	UUID           uuid.UUID `gorm:"index:idx_writing_uuid"`
	Title          string
	Tags           pq.StringArray `gorm:"type:text[];index:idx_writing_tags"`
	UserID         uint
	User           User
	CreatorID      uint
	Creator        Creator
	Rank           uint
	RelRank        uint
	Views          uint // weight 1
	Likes          uint // weight 50
	ListAdds       uint // weight 50
	LibAdds        uint // weight 200
	Donations      uint // weight 1000
	Flags          uint // weight -50
}

type Content struct {
	ID          uint      `gorm:"primaryKey"`
	UUID        uuid.UUID `gorm:"index:idx_content_uuid"`
	Title       string
	WritingID   uint
	Writing     Writing
	ContentType string
}

type Donation struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatorID   uint
	Creator     Creator
	DonatorID   uint
	Donator     User `gorm:"foreignKey:DonatorID"`
	RecipientID uint
	Recipient   User `gorm:"foreignKey:RecipientID"`
	WritingID   uint
	Writing     Writing
	Message     string
	Value uint
}

type FreecreateDonation struct {
	ID uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID uint
	User User
	Message string
	Value uint
}

type ReadWriting struct {
	UserID    uint `gorm:"primaryKey"`
	User      User
	WritingID uint `gorm:"primaryKey"`
	Writing   Writing
	CreatedAt time.Time
	UpdatedAt time.Time
	// creating composite primary key here ensures uniqueness / non-duplication and removes necessity for extraneous ID
}

type LikedWriting struct {
	UserID    uint `gorm:"primaryKey"`
	User      User
	WritingID uint `gorm:"primaryKey"`
	Writing   Writing
	CreatedAt time.Time
}

type ReadingListWriting struct {
	UserID    uint `gorm:"primaryKey"`
	User      User
	WritingID uint `gorm:"primaryKey"`
	Writing   Writing
	CreatedAt time.Time
}

type LibraryWriting struct {
	UserID    uint `gorm:"primaryKey"`
	User      User
	WritingID uint `gorm:"primaryKey"`
	Writing   Writing
	CreatedAt time.Time
}

type BookshelfWriting struct {
	BookshelfID uint `gorm:"primaryKey"`
	Bookshelf   Bookshelf
	WritingID   uint `gorm:"primaryKey"`
	Writing     Writing
	CreatedAt   time.Time
	Position    int
}
