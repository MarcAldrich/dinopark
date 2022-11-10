package dinopark

import "net/http"

type Error struct {
	code    int
	message string
}

func (e *Error) Error() string {
	return e.message
}

var (
	NotFound       = Error{http.StatusNotFound, "resource not found"}
	NotImplemented = Error{http.StatusNotImplemented, "not implemented"}
)
