package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_content_queries "freecreate/internal/db/pg_content/queries"
	pg_content_types "freecreate/internal/db/pg_content/types"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleCreateChapter(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, pgContentPools config.PgContentPools, pgContentQueries config.PgContentQueries, params pg_core_types.CreateChapterParams)(pg_core_types.CreatedChapter, *api_error.Error){
	createdChapter, createChapterErr := pg_core_queries.CreateChapter(ctx, pgCore, pgCoreQueries, params)

	createChapterContentParams := pg_content_types.CreateChapterContentParams{
		ChapterUUID: createdChapter.UUID,
	}

	createChapterContentErr := pg_content_queries.CreateChapterContent(ctx, createdChapter.ContentDb, pgContentPools, pgContentQueries, createChapterContentParams)
	if createChapterContentErr != nil {
		return createdChapter, createChapterContentErr
	}
	
	return createdChapter, createChapterErr
}