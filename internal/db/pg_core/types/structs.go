package pg_core_types

import "github.com/google/uuid"

// =========== USER STRUCTS ====================

type CreateUserParams struct {
	Email string
}

type CreatedUser struct {
	UserId int64
}

type UpdateUserInfoParams struct {
	Username       string
	UserHandle     string
	ReadingHistory bool
	IsAdult        bool
	UserId         int64
}

type UserInfo struct {
	Username       string
	UserHandle     string
	IsAdult        bool
	ReadingHistory bool
}

type DeleteUserParams struct {
	UserId int64
}

// ========== CREATOR STRUCTS ===================

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

type GetMyCreatorsParams struct {
	UserId int64
}

type MyCreatorsStruct struct {
	Name          string
	CreatorHandle string
	UUID          uuid.UUID
}

type GetMyCreatorParams struct {
	UserId int64
}

type MyCreator struct {
	Name string
	UUID uuid.UUID
}

type GetCreatorParams struct {
}

type Creator struct {
}
