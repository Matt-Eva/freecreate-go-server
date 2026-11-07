package cache

import (
	"fmt"
	"slices"
)

func genSearchKey(searchType string, writingType string, datePosted string, topics []string, tags []string)(string){

	searchKey := searchType + ":" + writingType + ":" + datePosted

	slices.Sort(topics)

	for _, val := range topics{
		searchKey = searchKey + ":" + val
	}

	slices.Sort(tags)

	for _, val := range tags {
		searchKey = searchKey + ":" + val
	}

	fmt.Println(searchKey)

	return searchKey
}