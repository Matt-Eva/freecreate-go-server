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
	ID                 uint `gorm:"primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	UUID               uuid.UUID `gorm:"index:idx_creator_uuid"`
	UserID             uint			`gorm:"index:idx_creator_name_user_id,unique"`
	User               User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name               string    `gorm:"not_null;index:idx_creator_name_user_id,unique"`
	Tags               pq.StringArray `gorm:"type:text[];index:idx_creator_tags"`
	Followers          uint
	Subscribers        uint
	RecurringDonations uint
	Donations          uint
	Views              uint
	Likes              uint
}

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not_null;uniqueIndex"`
}

// use composite primary key to avoid data redundancy
type ContentTag struct {
	ID        uint    `gorm:"primaryKey"`
	TagID     uint    `gorm:"index"`
	Tag       Tag     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WritingID uint    `gorm:"index"`
	Writing   Writing `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatorID uint    `gorm:"index"`
	Creator   Creator
	CreatedAt time.Time
}

type Writing struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	WritingType string    `gorm:"not_null"`
	UUID        uuid.UUID `gorm:"index:idx_writing_uuid"`
	Title       string
	Tags        pq.StringArray `gorm:"type:text[];index:idx_writing_tags"`
	UserID      uint
	User        User
	CreatorID   uint
	Creator     Creator
	Rank        uint
	RelRank     uint
	Views       uint // weight 1
	Likes       uint // weight 50
	ListAdds    uint // weight 50
	LibAdds     uint // weight 200
	Donations   uint // weight 1000
	Flags       uint // weight -50
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
	Value       uint
}

type FreecreateDonation struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	User      User
	Message   string
	Value     uint
}

type ReadWriting struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index"`
	User      User
	WritingID uint `gorm:"index"`
	Writing   Writing
	CreatedAt time.Time
	UpdatedAt time.Time
	// creating composite primary key here ensures uniqueness / non-duplication and removes necessity for extraneous ID
}

type LikedWriting struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index"`
	User      User
	WritingID uint `gorm:"index"`
	Writing   Writing
	CreatedAt time.Time
}

type ReadingListWriting struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index"`
	User      User
	WritingID uint `gorm:"index"`
	Writing   Writing
	CreatedAt time.Time
}

type LibraryWriting struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index"`
	User      User
	WritingID uint `gorm:"index"`
	Writing   Writing
	CreatedAt time.Time
}

type BookshelfWriting struct {
	ID          uint `gorm:"primaryKey"`
	BookshelfID uint `gorm:"index"`
	Bookshelf   Bookshelf
	WritingID   uint `gorm:"index"`
	Writing     Writing
	CreatedAt   time.Time
	Position    int
}
