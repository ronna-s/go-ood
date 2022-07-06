package animal

// Animal ...
type Animal interface {
	MakeSound() string
}

// Predator ...
type Predator interface {
	Eats(Animal) bool
	Eat(Animal)
}

type BaseAnimal struct {
	Sound string
}

// Deer ...
type Deer struct {
	BaseAnimal
}

func NewDeer() Deer {
	return Deer{}
}
