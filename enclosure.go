package dinopark

// This line ensures that the interface is fully implemented at compile time which avoids waiting for a runtime panic without it.
var _ EnclosureControl = (*Enclosure)(nil)

// The enclosure struct digitally represents the physical cage in the park
type Enclosure struct {
	contains   []*Dinosaur
	capacity   *EnclosureCapacity
	powerState bool
}

type EnclosureCapacity struct {
	species  string
	capacity uint16
}

type EnclosureControl interface {
	// Returns the dinos in the cage or error. If no dinosaurs are registered in the cage
	// an empty slice is returned.
	ListDinosInEnclosure(filterOpts *EnclosureFilter) (dinosInCage []*Dinosaur, err *Error)

	// SAFETY: Moving dinos is a safety-critical item as some loose dinos may pose a safety risk to
	// the parks Dino-Handler personnel as well as to the park infrastructure
	// Moving dinos between cages is an automated operation to protect park staff
	MoveDinosToEnclosure(dinosToMove []*Dinosaur) (moveSuccess bool, err *Error)

	// Sets the number of dinos of a specific species that can be in this enclosure.
	// WARNING: The dev team does not currently have a digital model to validate the number of
	// dinos of a speicies allowed to cohabitate per average space. Without this model we
	// rely entirely on the science team to accurately set the value.
	SetEnclosureCapacity(newEncCap *EnclosureCapacity) (configuredCapacity *EnclosureCapacity, err *Error)

	// Returns the current in-use enclosure settings
	ReadEnclosureState() (enclosureState *Enclosure, err *Error)

	// Allows the system to set the power state. Will return error if attempting
	// to power down an enclosure with dinos currently in the enclosure.
	CmdPwrState(commandedPwrState bool) (currentPwrState bool, err *Error)

	// Returns "ACTIVE" or "DOWN"
	GetPwrState() (currentPwrState string, err *Error)
}

type EnclosureFilter struct {
	//TODO: implement business reqs with filter statements over from readme
}

func (e *Enclosure) ListDinosInEnclosure(filterOpts *EnclosureFilter) (dinosInCage []*Dinosaur, err *Error) {
	return nil, &NotImplemented
}

func (e *Enclosure) MoveDinosToEnclosure(dinosToMove []*Dinosaur) (moveSuccess bool, err *Error) {
	return false, &NotImplemented
}

func (e *Enclosure) SetEnclosureCapacity(newEncCap *EnclosureCapacity) (configuredCapacity *EnclosureCapacity, err *Error) {
	//Validate input
	if newEncCap == nil {
		return nil, &EncInvalidConfig
	}

	//Safety Check: Enclosure can not currently be holding dinos
	if len(e.contains) != 0 {
		//ERROR: Can't change capacity if enc not empty
		return nil, &EncNotEmpty
	}

	//Update enclosure configuration
	e.capacity = newEncCap

	//Success-> return now-running capacity configuration
	return e.capacity, nil
}

func (e *Enclosure) ReadEnclosureState() (enclosureState *Enclosure, err *Error) {
	return e, nil
}

func (e *Enclosure) CmdPwrState(commandedPwrState bool) (currentPwrState bool, err *Error) {
	return false, &NotImplemented
}

func (e *Enclosure) GetPwrState() (currentPwrState string, err *Error) {
	return "", &NotImplemented
}
