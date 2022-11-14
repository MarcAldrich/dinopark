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
	NotFound         = Error{http.StatusNotFound, "resource not found"}
	NotImplemented   = Error{http.StatusNotImplemented, "not implemented"}
	EncInvalidConfig = Error{http.StatusUnprocessableEntity, "invalid enclosure configuration requested"}
	EncNotEmpty      = Error{http.StatusConflict, "unable to adjust enclosure settings while dinos in the enclosure"}
	MissingArg       = Error{http.StatusBadRequest, "missing an argument"}
)
