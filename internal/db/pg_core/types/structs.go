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
	ID            int
	UUID          uuid.UUID
}

type GetMyCreatorParams struct {
	UserId int64
}

type MyCreator struct {
	Name string
	ID   int64
	UUID uuid.UUID
}

type GetCreatorParams struct {
}

type Creator struct {
}

// ========= Writing Structs =======

type NewWritingParams struct {
	Title       string
	CreatorId   int64
	UserId      int64
	WritingType string
}

type CreatedWriting struct {
	UUID        uuid.UUID
	UserId      int64
	CreatorId   int64
	Title       string
	Subtitle    string
	WritingType string
	Topics      []string
	Tags        []string
	Description string
	Published   bool
}

type GetMyWritingsParams struct {
	UserId int64
}

type MyWritingsStruct struct {
	UUID      uuid.UUID
	CreatorId int64
	Title     string
	Published bool
}
