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
	ID            int64
	UUID          uuid.UUID
}

type GetMyCreatorParams struct {
	UserId      int64
	CreatorUUID uuid.UUID
}

type MyCreator struct {
	Name string
	ID   int64
	UUID uuid.UUID
}

type GetCreatorParams struct {
	UUID uuid.UUID
}

type Creator struct {
	Name string
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

type GetMyWritingParams struct {
	UUID   uuid.UUID
	UserId int64
}

type MyWriting struct {
	ID          int64
	UUID        uuid.UUID
	Title       string
	Subtitle    string
	Description string
	Topics      []string
	Tags        []string
	WritingType string
	Published   bool
	ListAdds    int64
	Likes       int64
	LibAdds     int64
	Donations   int64
	Flags       int64
	IsAdult 	bool
	CreatorId   int64
	CreatorUUID uuid.UUID
	CreatorName string
}

type GetMyCreatorWritingsParams struct {
	UserId    int64
	CreatorId int64
}

type MyCreatorWritings struct {
	Title       string
	UUID        uuid.UUID
	WritingType string
	Published   bool
}

type GetWritingParams struct {
	UUID uuid.UUID
}

type Writing struct {
	UUID        uuid.UUID
	Title       string
	Subtitle    string
	Description string
	Topics      []string
	Tags        []string
}

// ===== Chapter Structs =====

type CreateChapterParams struct {
	Title         string
	ChapterNumber int
}

type CreatedChapter struct {
	Title         string
	ChapterNumber int
	Published     bool
	UUID          uuid.UUID
}

type GetMyWritingChapterParams struct {
	UserId    int64
	WritingId int64
}

type GetMyChapterParams struct {
	UserId int64
	UUID   uuid.UUID
}

type GetChapterParams struct {
	UUID uuid.UUID
}

type Chapter struct {
	UUID          uuid.UUID
	Title         string
	ChapterNumber int
	Published     bool
}
