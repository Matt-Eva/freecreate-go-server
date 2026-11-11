package pg_content_types

import "github.com/google/uuid"

type CreateChapterContentParams struct {
	ChapterUUID uuid.UUID
	UserUuid uuid.UUID
}

type GetJsonChapterContentParams struct {
	ChapterUuid uuid.UUID
	ContentDB int
}

type JsonChapterContent struct {
	JsonContent map[string]any
}
