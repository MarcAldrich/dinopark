package dinopark

// This line ensures that the interface is fully implemented at compile time which avoids waiting for a runtime panic without it.
var _ ParkControl = (*Park)(nil)

type Park struct {
	Dinos      []Dinosaur  //Holds references to all dinos in the park
	Enclosures []Enclosure //Holds references to all enclosures in park
	Places     []Place     //Holds references to all places in a park

}

func NewPark() *Park {
	return &Park{
		Dinos:      []Dinosaur{},
		Enclosures: []Enclosure{},
		Places:     []Place{},
	}
}

type ParkControl interface {
	//PLACE CONTROL
	//List places {by-type, by-type-and-power-status, ...}
	ListPlaces(filter *PlaceFilter) (places []Place, err *Error)
	//Register a new enclosure
	RegisterPlace(plcToReg *Enclosure) (plcAdded *Enclosure, err *Error)
	//Deregisteres a decomissioned place {lab, enclosure, etc}
	RemovePlace(plcToRemove *Enclosure) (plcRemoved *Place, err *Error)

	//DINO CONTROL
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

func (p Park) ListPlaces(filter *PlaceFilter) (places []Place, err *Error) {
	//Short-circuit: If filter empty return all
	if filter == nil {
		return p.Places, nil
	}

	//Was a filter specified?
	var plTypeFilter *PlaceKind
	if filter.ByKind != nil {
		plTypeFilter = filter.ByKind
	}

	//Build result by filter
	for _, pl := range p.Places {
		if plTypeFilter == &pl.Kind {
			//Case: Filter matches entry add to result
			places = append(places, pl)
		}
	}

	return places, nil
}

func (p *Park) RegisterPlace(plcToReg *Enclosure) (plcAdded *Enclosure, err *Error) {
	return nil, &NotImplemented
}

func (p *Park) RemovePlace(plcToRemove *Enclosure) (plcRemoved *Place, err *Error) {
	return nil, &NotImplemented
}

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
