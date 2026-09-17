package pg_core_types

import "github.com/google/uuid"

type UserInfo struct {
	Username       string
	UserHandle     string
	IsAdult        bool
	ReadingHistory bool
}

type UpdateUserInfoParams struct {
	Username       string `db:"username"`
	UserHandle     string `db:"user_handle"`
	ReadingHistory bool   `db:"reading_history"`
	IsAdult        bool   `db:"is_adult"`
	UserId         int64  `db:"user_id"`
}

type CreateCreatorParams struct {
	UserId        int64 `db:"user_id"`
	Name          string
	CreatorHandle string `db:"creator_handle"`
}

type CreatedCreator struct {
	Name          string    `db:"name"`
	CreatorHandle string    `db:"creator_handle"`
	UUID          uuid.UUID `db:"uuid"`
}
