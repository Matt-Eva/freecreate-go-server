package api_error

type Error struct {
	Code int
	Message string
	Error error
}

var InteralServerErrorMessage = "There was an issue processing that request."