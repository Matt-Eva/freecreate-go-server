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

func HandleGetEditChapter(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, pgContentPools config.PgContentPools, pgContentQueries config.PgContentQueries, userId int64, writingUuid uuid.UUID, chapterUuid uuid.UUID)(pg_core_types.MyWriting, pg_core_types.Chapter, pg_content_types.JsonChapterContent, *api_error.Error){
	var myWriting pg_core_types.MyWriting
	var myChapter pg_core_types.Chapter
	var jsonChapterContent pg_content_types.JsonChapterContent

	getMyWritingParams := pg_core_types.GetMyWritingParams{
		UUID: writingUuid,
		UserId: userId,
	}

	myWriting, getMyWritingErr := pg_core_queries.GetMyWriting(ctx, pgCore, pgCoreQueries, getMyWritingParams)
	if getMyWritingErr != nil {
		return myWriting, myChapter, jsonChapterContent, getMyWritingErr
	}

	getMyChapterParams := pg_core_types.GetMyChapterParams{
		UserId: userId,
		UUID: chapterUuid,
	}

	myChapter, getMyChapterErr := pg_core_queries.GetMyChapter(ctx, pgCore, pgCoreQueries, getMyChapterParams)
	if getMyChapterErr != nil {
		return myWriting, myChapter, jsonChapterContent, getMyChapterErr
	}

	getJsonChapterContentParams := pg_content_types.GetJsonChapterContentParams{
		ChapterUuid: chapterUuid,
		ContentDB: int(myChapter.ContentDb),
	}

	jsonChapterContent, getJsonChapterContentErr := pg_content_queries.GetJsonChapterContent(ctx, pgContentPools, pgContentQueries, getJsonChapterContentParams)
	if getJsonChapterContentErr != nil {
		return myWriting, myChapter, jsonChapterContent, getJsonChapterContentErr
	}

	return myWriting, myChapter, jsonChapterContent, nil
}