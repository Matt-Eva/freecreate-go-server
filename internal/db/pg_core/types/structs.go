package pg_core_types

import "github.com/google/uuid"

type UserInfo struct {
	Username       string
	UserHandle     string
	IsAdult        bool
	ReadingHistory bool
}

type UpdateUserInfoParams struct {
	Username       string
	UserHandle     string
	ReadingHistory bool
	IsAdult        bool
	UserId         int64
}

type NewCreatorParams struct {
	UserId        int64
	Name          string
	CreatorHandle string
}

type CreatedCreator struct {
	Name          string
	CreatorHandle string
	UUID          uuid.UUID
}
