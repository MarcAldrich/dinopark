package dinopark

// This line ensures that the interface is fully implemented at compile time which avoids waiting for a runtime panic without it.
var _ ParkControl = (*Park)(nil)

type Park struct {
	Dinos      []Dinosaur  //Holds references to all dinos in the park
	Enclosures []Enclosure //Holds references to all enclosures in park
	Places     *Places     //Holds references to all places in a park
}

func NewPark() *Park {
	return &Park{
		Dinos:      []Dinosaur{},
		Enclosures: []Enclosure{},
		Places:     NewPlaces(), //NOTE: Use new func to ensure mutex is initalized
	}
}

type ParkControl interface {
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

func (p *Park) RemovePlace(plcToRemove *Place) (plcRemoved *Place, err *Error) {
	//Error check: missing arg
	if plcToRemove == nil {
		return nil, &MissingArg
	}

	//Validate input
	if !plcToRemove.Validate() {
		return nil, &InvalidArg
	}

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
