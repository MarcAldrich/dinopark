package dinopark

// This line ensures that the interface is fully implemented at compile time which avoids waiting for a runtime panic without it.
var _ ParkControl = (*Park)(nil)

type Park struct {
	Places []Place    //Holds all places in a park
	Dinos  []Dinosaur //Holds all dinos in the park
}

type ParkControl interface {
	//PLACE CONTROL
	//List places {by-type, by-type-and-power-status, ...}
	ListPlaces(filter PlaceFilter) (places []Place, err *Error)
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

func (p Park) ListPlaces(filter PlaceFilter) (places []Place, err *Error) {
	return nil, &NotImplemented
}

func (p Park) RegisterPlace(plcToReg *Enclosure) (plcAdded *Enclosure, err *Error) {
	return nil, &NotImplemented
}

func (p Park) RemovePlace(plcToRemove *Enclosure) (plcRemoved *Place, err *Error) {
	return nil, &NotImplemented
}

func (p Park) ListDinos(filter DinoFilter) (dinos []Dinosaur, err *Error) {
	return nil, &NotImplemented
}

func (p Park) RegisterDino(dinoToReg *Dinosaur) (dinoReged *Dinosaur, err *Error) {
	return nil, &NotImplemented
}

func (p Park) RemoveDino(dinoToRem *Dinosaur) (dinoRmed *Dinosaur, err *Error) {
	return nil, &NotImplemented
}

func (p Park) MoveDinos(dinosToMove []*Dinosaur, dstPlc *Place) (dinosMoved []*Dinosaur, err *Error) {
	return nil, &NotImplemented
}
