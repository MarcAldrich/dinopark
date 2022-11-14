package dinopark

import (
	"github.com/google/uuid"
)

type Dinosaur struct {
	ID      uuid.UUID
	Name    string
	Species DinoSpecies
}

func NewDinosaur(name string, species DinoSpecies) (dino *Dinosaur, err *Error) {
	if name == "" {
		return nil, &MissingArg
	}

	return &Dinosaur{
		ID:      uuid.New(),
		Name:    name,
		Species: species,
	}, err
}

type DinoFilter struct {
	ByDiet    *DinoDiet
	BySpecies *DinoSpecies
	ByName    *string
}

type DinoSpecies uint16

const (
	Tyrannosaurus = iota
	Velociraptor
	Spinosaurus
	Megalosaurus
	Brachiosaurus
	Stegosaurus
	Ankylosaurus
	Triceratops
)

type DinoDiet uint16

const (
	Omnivore = iota
	Carnivore
	Herbivore
)

func (d *Dinosaur) DinoDietType() DinoDiet {
	var dDiet DinoDiet
	switch d.Species {
	case Tyrannosaurus:
		dDiet = Carnivore
	case Velociraptor:
		dDiet = Carnivore
	case Spinosaurus:
		dDiet = Carnivore
	case Megalosaurus:
		dDiet = Carnivore
	case Brachiosaurus:
		dDiet = Herbivore
	case Stegosaurus:
		dDiet = Herbivore
	case Ankylosaurus:
		dDiet = Herbivore
	case Triceratops:
		dDiet = Herbivore
	}

	return dDiet
}
