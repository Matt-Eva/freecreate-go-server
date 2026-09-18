package pg_core_types

import "github.com/google/uuid"

type CreateUserParams struct {
	Email string
}

type CreatedUser struct {
	UserId int64
}

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

type MyCreatorsStruct struct {
	Name          string
	CreatorHandle string
	UUID          uuid.UUID
}

type MyCreator struct {
	Name          string
	UUID          uuid.UUID
}

type Creator struct {
}
