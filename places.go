package dinopark

import (
	"sync"

	"github.com/google/uuid"
)

//Concurrency-safe data structure via mutex lock around slice
type Places struct {
	Places map[uuid.UUID]*Place
	mu     *sync.Mutex
}

func NewPlaces() *Places {
	return &Places{
		Places: map[uuid.UUID]*Place{},
		mu:     &sync.Mutex{},
	}
}

type PlaceControl interface {
	//List places {by-type, by-type-and-power-status, ...}
	ListPlaces(filter *PlaceFilter) (places []Place, err *Error)
	//Register a new place
	RegisterPlace(plcToReg *Place) (plcAdded *Place, err *Error)
	//Deregisters a decomissioned place {lab, enclosure, etc}
	RemovePlace(plcToRemove *Place) (plcRemoved *Place, err *Error)
}

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
	ByKind *PlaceKind
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

//
// PLACE CONTROL
//

func (pls Places) ListPlaces(filter *PlaceFilter) (places []*Place, err *Error) {
	//Short-circuit: If filter empty return all
	if filter == nil {
		return pls.PlacesToSlice(), nil
	}

	//Was a filter specified?
	var plTypeFilter *PlaceKind
	if filter.ByKind != nil {
		plTypeFilter = filter.ByKind
	}

	//Build result by filter
	for _, pl := range pls.Places {
		if plTypeFilter == &pl.Kind {
			//Case: Filter matches entry add to result
			places = append(places, pl)
		}
	}

	return places, nil
}

// Helper func to output slice from map: Referenced from maps.Values implementation in the golang.org/x/exp
func (pls Places) PlacesToSlice() []*Place {
	r := make([]*Place, 0, len(pls.Places))
	for _, v := range pls.Places {
		r = append(r, v)
	}
	return r
}

func (pls *Places) RegisterPlace(plcToReg *Place) (plcAdded *Place, err *Error) {
	//Error check: missing arg
	if plcToReg == nil {
		return nil, &MissingArg
	}

	//Validate input
	if !plcToReg.Validate() {
		return nil, &InvalidArg
	}

	//CONCURRENT SECTION
	pls.mu.Lock()
	defer pls.mu.Unlock()
	//HOLDING LOCK
	pls.Places[plcToReg.ID] = plcToReg

	return plcToReg, nil
	//LOCK RELEASE ON FUNC EXIT VIA DEFER
}

func (pls *Places) RemovePlace(plcToRemove *Place) (plcRemoved *Place, err *Error) {
	//Error check: missing arg
	if plcToRemove == nil {
		return nil, &MissingArg
	}

	//Validate input
	if !plcToRemove.Validate() {
		return nil, &InvalidArg
	}

	mapLenBeforeDelete := len(pls.Places)
	//CONCURRENT SECTION
	pls.mu.Lock()
	pls.mu.Unlock()
	//HOLDING LOCK: Critical section
	delete(pls.Places, plcToRemove.ID) //delete is a no-op if not found; use len of map as validation of delete
	if mapLenBeforeDelete-len(pls.Places) != 1 {
		//Error: entry not found
		return nil, &NotFound
	}

	return plcToRemove, nil
}
