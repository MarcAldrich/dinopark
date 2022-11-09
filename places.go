package dinopark

//
type places struct {
	name     string
	location string
	kind     PlaceType
}

type PlaceType uint8

const (
	Lab PlaceType = iota
	Cage
)
