package dinopark

type Dinosaur struct {
	Name    string
	Species DinoSpecies
}

type DinoFilter struct {
	ByDiet    string
	BySpecies string
	ByName    string
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
