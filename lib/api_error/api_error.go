package api_error

type Error struct {
	Code int
	Message string
	Error error
}