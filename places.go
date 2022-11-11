package dinopark

type Place struct {
	Name     string
	Location string
	Kind     PlaceType
}

type PlaceFilter struct{}

type PlaceType uint8

const (
	LAB PlaceType = iota
	ENCLOSURE
)

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
