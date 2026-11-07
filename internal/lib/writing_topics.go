package lib

import (
	"slices"
)

type WritingTopicsStruct struct {
	EssaysAndBlogs         []string
	Fiction                []string
	MemoirAndAutobiography []string
	Poetry                 []string
}

var WritingTopics = WritingTopicsStruct{
	EssaysAndBlogs:         []string{"No Topic", "Culture", "Economics", "Politics", "Technology"},
	Fiction:                []string{"No Genre", "Action", "Adventure", "Comedy", "Drama", "Epic", "Erotica", "Fantasy", "Historical Fiction", "Horror", "Literary Fiction", "Magical Realism", "Mystery", "Realism", "Romance", "Science Fiction", "Speculative Fiction", "Social Fiction", "Superhero", "Supernatural", "Thriller"},
	MemoirAndAutobiography: []string{"No Topic", "Travel"},
	Poetry:                 []string{"No Genre", "Long Form Poetry", "Short Form Poetry"},
}

func generateAllWritingTopics() []string{
	var allWritingTopics []string

	for _, val := range WritingTopics.Fiction {
			allWritingTopics = append(allWritingTopics, val)
		}

		for _, val := range WritingTopics.EssaysAndBlogs {
			allWritingTopics = append(allWritingTopics, val)
		}

		for _, val := range WritingTopics.MemoirAndAutobiography {
			allWritingTopics = append(allWritingTopics, val)
		}

		for _, val := range WritingTopics.Poetry {
			allWritingTopics = append(allWritingTopics, val)
		}

		slices.Sort(allWritingTopics)
		allWritingTopics = slices.Compact(allWritingTopics)
		allWritingTopics = slices.DeleteFunc(allWritingTopics, func(topic string) bool {
			return topic == "No Genre" || topic == "No Topic"
		})

		return allWritingTopics
}

var AllWritingTopics = generateAllWritingTopics()