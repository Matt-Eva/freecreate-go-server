package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_content_queries "freecreate/internal/db/pg_content/queries"
	pg_content_types "freecreate/internal/db/pg_content/types"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HandleCreateChapterParams struct {
	WritingUUID uuid.UUID
	UserId int64
	Title string
	ChapterNumber int
}

func HandleCreateChapter(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, pgContentPools config.PgContentPools, pgContentQueries config.PgContentQueries, params HandleCreateChapterParams)(pg_core_types.CreatedChapter, *api_error.Error){

	getMyWritingParams := pg_core_types.GetMyWritingParams{
		UUID: params.WritingUUID,
		UserId: params.UserId,
	}

	myWriting, getMyWritingErr := pg_core_queries.GetMyWriting(ctx, pgCore, pgCoreQueries, getMyWritingParams)
	if getMyWritingErr != nil {
		return pg_core_types.CreatedChapter{}, getMyWritingErr
	}

	createChapterParams := pg_core_types.CreateChapterParams{
		UserId: params.UserId,
		WritingId: myWriting.ID,
		CreatorId: myWriting.CreatorId,
		Title: params.Title,
		ChapterNumber: params.ChapterNumber,
	}

	createdChapter, createChapterErr := pg_core_queries.CreateChapter(ctx, pgCore, pgCoreQueries, createChapterParams)
	if createChapterErr != nil {
		return createdChapter, createChapterErr
	}

	createChapterContentParams := pg_content_types.CreateChapterContentParams{
		ChapterUUID: createdChapter.UUID,
	}

	createChapterContentErr := pg_content_queries.CreateChapterContent(ctx, createdChapter.ContentDb, pgContentPools, pgContentQueries, createChapterContentParams)
	if createChapterContentErr != nil {
		return createdChapter, createChapterContentErr
	}
	
	return createdChapter, createChapterErr
}