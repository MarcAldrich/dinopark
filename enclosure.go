package dinopark

import "github.com/google/uuid"

// This line ensures that the interface is fully implemented at compile time which avoids waiting for a runtime panic without it.
var _ EnclosureControl = (*Enclosure)(nil)

// The enclosure struct digitally represents the physical cage in the park
type Enclosure struct {
	ID         uuid.UUID
	contains   []*Dinosaur
	capacity   *EnclosureCapacity
	powerState bool
}

func NewEnclosure(cap *EnclosureCapacity) (newEnc *Enclosure, err *Error) {
	if cap == nil {
		return nil, &MissingArg
	}
	return &Enclosure{
		ID:         uuid.New(),
		contains:   []*Dinosaur{},
		capacity:   cap,
		powerState: false,
	}, err
}

type EnclosureCapacity struct {
	species  string
	capacity uint16
}

type EnclosureControl interface {
	// Returns the dinos in the cage or error. If no dinosaurs are registered in the cage
	// an empty slice is returned.
	ListDinosInEnclosure(filterOpts *EnclosureFilter) (dinosInCage []*Dinosaur, err *Error)

	// Sets the number of dinos of a specific species that can be in this enclosure.
	// WARNING: The dev team does not currently have a digital model to validate the number of
	// dinos of a speicies allowed to cohabitate per average space. Without this model we
	// rely entirely on the science team to accurately set the value.
	SetEnclosureCapacity(newEncCap *EnclosureCapacity) (configuredCapacity *EnclosureCapacity, err *Error)

	// Returns the current in-use enclosure settings
	ReadEnclosureState() (enclosureState *Enclosure, err *Error)

	// Allows the system to set the power state. Will return error if attempting
	// to power down an enclosure with dinos currently in the enclosure.
	CmdPwrState(commandedPwrState bool) (currentPwrState *bool, err *Error)

	// Returns "ACTIVE" or "DOWN"
	GetPwrState() (currentPwrState string)
}

type EnclosureFilter struct {
	ByPowerState        *bool        //Use-case: Find all "DOWN" enclosures
	ByNumberOfDinos     *uint16      //Helpful to find empty enclosures
	ByCapacityRemaining *uint16      //Use-case: find an enclosure with room for a dino
	BySpecies           *DinoSpecies //Use-case: find enclosures with same species
	ByDiet              *DinoDiet    //Use-case: find an enclosure with only herbivores/carnivores
}

func (e *Enclosure) ListDinosInEnclosure(filterOpts *EnclosureFilter) (dinosInCage []*Dinosaur, err *Error) {

	return nil, &NotImplemented
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
	//NOTE: Short implementation: This abstraction provides an easy place to add any data/config validation required in the future
	return e, nil
}

func (e *Enclosure) CmdPwrState(commandedPwrState bool) (currentPwrState *bool, err *Error) {
	//Validate/Optimization: Do nothing if commanded state matches existing power state
	if commandedPwrState == e.powerState {
		return &e.powerState, nil
	}

	//Safety Check: Do not power down if dinos in enclosure
	if len(e.contains) > 0 && !commandedPwrState {
		return nil, &EncNotEmpty
	}

	//apply power state
	e.powerState = commandedPwrState

	return &e.powerState, nil
}

func (e *Enclosure) GetPwrState() (currentPwrState string) {
	if e.powerState {
		return "ACTIVE"
	}

	return "DOWN"
}
