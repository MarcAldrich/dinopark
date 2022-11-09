package dinopark

// The enclosure struct digitally represents the physical cage in the park
type Enclosure struct {
	contains   []*Dinosaur
	capacity   uint16
	powerState bool
}

type EnclosureControl interface {
	// Returns the dinos in the cage or error. If no dinosaurs are registered in the cage
	// an empty slice is returned.
	ListDinosInCage(filterOpts *EnclosureFilter) (dinosInCage []*Dinosaur, err error)

	// SAFETY: Moving dinos is a safety-critical item as some loose dinos may pose a safety risk to
	// the parks Dino-Handler personnel as well as to the park infrastructure
	//
	RequestMoveDinosToCage(dinosToMove []*Dinosaur) (moveSuccess bool, err error)

	// Sets the number of dinos of a specific species that can be in this enclosure.
	// WARNING: The dev team does not currently have a digital model to validate the number of
	// dinos of a speicies allowed to cohabitate per average space. Without this model we
	// rely entirely on the science team to accurately set the value.
	SetEnclosureCapacity(dinoSpecies string) (err error) //TODO: Finish this protoype by updating the dinoSpecies type
}

type EnclosureFilter struct {
	//TODO: implement business reqs with filter statements over from readme
}
