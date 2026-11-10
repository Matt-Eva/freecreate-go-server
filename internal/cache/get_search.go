package cache

import (
	"context"

	"github.com/valkey-io/valkey-go"
)

func GetSearch(ctx context.Context, valkeyClient valkey.Client, searchType string, writingType string, datePosted string, topics []string, tags []string){
	genSearchKey(searchType, writingType, datePosted, topics, tags)
}