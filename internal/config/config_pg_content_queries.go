package config

import (
	"freecreate/internal/lib/logger"
	"os"
	"path/filepath"
)

type PgContentQueries struct {
	createChapterContent string
	getJsonChapterContent string
}

func (q PgContentQueries) CreateChapterContent() string {
	return q.createChapterContent
}

func (q PgContentQueries) GetJsonChapterContent() string{
	return q.getJsonChapterContent
}

func ConfigPgContentQueries() (PgContentQueries, error) {
	q := PgContentQueries{}

	createChapterContent, createChapterContentError := os.ReadFile(filepath.Join("./internal/db/pg_content/query_files", "create_chapter_content.sql"))
	if createChapterContentError != nil {
		logger.Log(createChapterContentError)
		return q, createChapterContentError
	}
	q.createChapterContent = string(createChapterContent)

	getJsonChapterContent, getJsonChapterContentError := os.ReadFile(filepath.Join("./internal/db/pg_content/query_files", "get_json_chapter_content.sql"))
	if getJsonChapterContentError != nil {
		logger.Log(getJsonChapterContentError)
		return q, getJsonChapterContentError
	}
	q.getJsonChapterContent = string(getJsonChapterContent)

	return q, nil
}
