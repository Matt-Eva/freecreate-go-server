package config

import (
	"freecreate/internal/lib/logger"
	"os"
	"path/filepath"
)

type PgContentQueries struct {
	createChapterContent string
}

func (q PgContentQueries) CreateChapterContent() string {
	return q.createChapterContent
}

func ConfigPgContentQueries() (PgContentQueries, error) {
	q := PgContentQueries{}

	createChapterContent, createChapterContentError := os.ReadFile(filepath.Join("./internal/db/pg_content/query_files", "create_chapter_content.sql"))
	if createChapterContentError != nil {
		logger.Log(createChapterContentError)
		return q, createChapterContentError
	}
	q.createChapterContent = string(createChapterContent)

	return q, nil
}
