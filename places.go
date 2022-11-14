package dinopark

import "github.com/google/uuid"

type Place struct {
	ID       uuid.UUID
	Name     string
	Location string
	Kind     PlaceKind
}

func NewPlace(name string, loc string, kind PlaceKind) (pl *Place, err *Error) {
	if name == "" {
		return nil, &MissingArg
	}

	if loc == "" {
		return nil, &MissingArg
	}

	return &Place{
		ID:       uuid.New(),
		Name:     name,
		Location: loc,
		Kind:     kind,
	}, err
}

type PlaceFilter struct {
	ByKind *PlaceKind `json:'by_kind'`
}

type PlaceKind uint8

const (
	LAB PlaceKind = iota
	ENCLOSURE
)

func NewPlaceKind(kind PlaceKind) *PlaceKind {
	plKind := kind
	return &plKind
}

func (p Place) Validate() (isValid bool) {
	if p.Name == "" {
		return false
	}

	if p.Location == "" {
		return false
	}

	return true
}

func (p Place) String() string {
	strToRet := ""
	switch p.Kind {
	case LAB:
		strToRet = "Lab"
	case ENCLOSURE:
		strToRet = "Enclosure"
	}
	return strToRet
}
