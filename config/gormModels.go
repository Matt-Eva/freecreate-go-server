package config

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"not_null;uniqueIndex"`
}
