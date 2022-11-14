package dinopark

type Place struct {
	Name     string
	Location string
	Kind     PlaceKind
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
