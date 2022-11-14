package dinopark

import (
	"sync"
)

// This line ensures that the interface is fully implemented at compile time which avoids waiting for a runtime panic without it.
var _ ParkControl = (*Park)(nil)

type Park struct {
	Dinos      []Dinosaur  //Holds references to all dinos in the park
	Enclosures []Enclosure //Holds references to all enclosures in park
	Places     Places      //Holds references to all places in a park
}

//Concurrency-safe data structure via mutex lock around slice
type Places struct {
	Places []Place
	mu     *sync.Mutex
}

func NewPark() *Park {
	return &Park{
		Dinos:      []Dinosaur{},
		Enclosures: []Enclosure{},
		Places:     Places{},
	}
}

type ParkControl interface {
	//
	//PLACE CONTROL
	//
	//List places {by-type, by-type-and-power-status, ...}
	ListPlaces(filter *PlaceFilter) (places []Place, err *Error)
	//Register a new place
	RegisterPlace(plcToReg *Place) (plcAdded *Place, err *Error)
	//Deregisters a decomissioned place {lab, enclosure, etc}
	RemovePlace(plcToRemove *Place) (plcRemoved *Place, err *Error)

	//
	//ENCLOSURE CONTROL
	//
	//List enclosures with optional EnclosureFilter
	ListEnclosures(filter *EnclosureFilter) (encs []Enclosure, err *Error)
	//Register a new enclosure
	RegisterEnclosure(encToReg *Enclosure) (encAdded *Enclosure, err *Error)
	//Deregisters a decomissioned enclosure
	RemoveEnclosure(encToRem *Enclosure) (encRemoved *Enclosure, err *Error)

	//
	//DINO CONTROL
	//
	//List dinosaurs using dino filter
	ListDinos(filter DinoFilter) (dinos []Dinosaur, err *Error)
	//Registers a new dino
	RegisterDino(dinoToReg *Dinosaur) (dinoReged *Dinosaur, err *Error)
	//Removes a registered dino
	//NOTE: useful after events like that fatal velociraptor move from lab to enclosure. Never forget: "Shooooooot her"
	//RIP Jophery Brown (https://listofdeaths.fandom.com/wiki/Jurassic_Park#Jurassic_Park)
	RemoveDino(dinoToRem *Dinosaur) (dinoRmed *Dinosaur, err *Error)
	// SAFETY: Moving dinos is a safety-critical item as some loose dinos may pose a safety risk to
	// the parks Dino-Handler personnel as well as to the park infrastructure
	// Moving dinos between places is an automated operation to protect park staff
	MoveDinos(dinosToMove []*Dinosaur, dstPlc *Place) (dinosMoved []*Dinosaur, err *Error)
}

//
// PLACE CONTROL
//

func (p Park) ListPlaces(filter *PlaceFilter) (places []Place, err *Error) {
	//Short-circuit: If filter empty return all
	if filter == nil {
		return p.Places.Places, nil
	}

	//Was a filter specified?
	var plTypeFilter *PlaceKind
	if filter.ByKind != nil {
		plTypeFilter = filter.ByKind
	}

	//Build result by filter
	for _, pl := range p.Places.Places {
		if plTypeFilter == &pl.Kind {
			//Case: Filter matches entry add to result
			places = append(places, pl)
		}
	}

	return places, nil
}

func (p *Park) RegisterPlace(plcToReg *Place) (plcAdded *Place, err *Error) {
	//Error: missing arg
	if plcToReg == nil {
		return nil, &MissingArg
	}

	//Validate input
	if !plcToReg.Validate() {
		return nil, &MissingArg
	}

	//CONCURRENT SECTION
	p.Places.mu.Lock()
	defer p.Places.mu.Unlock()
	//HOLDING LOCK
	p.Places.Places = append(p.Places.Places, *plcToReg)

	return plcToReg, nil
	//LOCK RELEASE ON FUNC EXIT
}

func (p *Park) RemovePlace(plcToRemove *Place) (plcRemoved *Place, err *Error) {
	return nil, &NotImplemented
}

//
// ENCLOSURE CONTROL
//
func (p Park) ListEnclosures(filter *EnclosureFilter) (encs []Enclosure, err *Error) {
	return nil, &NotImplemented
}

func (p *Park) RegisterEnclosure(encToReg *Enclosure) (encAdded *Enclosure, err *Error) {
	return nil, &NotImplemented
}

func (p *Park) RemoveEnclosure(encToRem *Enclosure) (encRemoved *Enclosure, err *Error) {
	return nil, &NotImplemented
}

//
// DINO CONTROL
//

func (p Park) ListDinos(filter DinoFilter) (dinos []Dinosaur, err *Error) {
	return nil, &NotImplemented
}

func (p *Park) RegisterDino(dinoToReg *Dinosaur) (dinoReged *Dinosaur, err *Error) {
	return nil, &NotImplemented
}

func (p *Park) RemoveDino(dinoToRem *Dinosaur) (dinoRmed *Dinosaur, err *Error) {
	return nil, &NotImplemented
}

func (p Park) MoveDinos(dinosToMove []*Dinosaur, dstPlc *Place) (dinosMoved []*Dinosaur, err *Error) {
	return nil, &NotImplemented
}
