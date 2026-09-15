package query_handlers

import "freecreate/internal/lib/api_error"

type CreateCreatorParams struct {
	Name string
	Handle string
}

type CreatedCreator struct {
	Name string
	Handle string
	UUID string
}

func HandleCreateCreator()(CreatedCreator, *api_error.Error){
	
}