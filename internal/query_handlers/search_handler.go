package query_handlers

import (
	"context"
	"freecreate/internal/cache"
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

type SearchResults struct {
	Topics []string
	Tags	[]string
}

func HandleSearch(ctx context.Context, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, searchType string, writingType string, datePosted string, topics []string, tags []string)(SearchResults, *api_error.Error){

	var truncatedTopics []string
	
	for _, val := range topics {
		if len(truncatedTopics) < 3{
			if len(topics) > 1 && (val == "No Topic" || val == "No Genre"){
				continue
			}
			truncatedTopics = append(truncatedTopics, val)
		} else {
			break
		}
	}

	var truncatedTags []string

	for _, val := range tags {
		if len(tags) < 5 {
			truncatedTags = append(truncatedTags, val)
		} else {
			break
		}
	}

	cache.GetSearch(ctx, valkeyClient, searchType, writingType, datePosted, topics, tags)

	searchResults := SearchResults{
		Topics: truncatedTopics,
		Tags: truncatedTags,
	}

	return searchResults, nil
}