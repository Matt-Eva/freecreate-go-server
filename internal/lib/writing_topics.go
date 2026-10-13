package lib

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
	Poetry:                 []string{"No Genre", "Long Form", "Short Form"},
}
