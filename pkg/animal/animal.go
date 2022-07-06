package animal

import "time"

// Now ...
var Now = time.Now

// Animal ...
type Animal interface {
	MakeSound() string
}

// Predator ...
type Predator interface {
	Animal
	Eats(a Animal) bool
	Eat(a Animal) string
	Full() bool
}

// BaseAnimal ...
type BaseAnimal struct {
	Sound string
}

// MakeSound ...
func (b BaseAnimal) MakeSound() string {
	return b.Sound
}

// NewDeer ...
func NewDeer() Deer {
	return Deer{BaseAnimal{"I'm so cute!"}}
}

// Deer ...
type Deer struct {
	BaseAnimal
}

// NewLion ...
func NewLion() *Lion {
	return &Lion{BaseAnimal: BaseAnimal{Sound: "Roar!"}}
}

// Lion ...
type Lion struct {
	BaseAnimal
	LastFed time.Time
}

// Eats ...
func (l Lion) Eats(a Animal) bool {
	switch a.(type) {
	case Deer:
		return true
	}
	return false
}

// Eat ...
func (l *Lion) Eat(_ Animal) string {
	l.LastFed = Now()
	return l.MakeSound()
}

// Full ...
func (l Lion) Full() bool {
	return l.LastFed.After(Now().Add(time.Minute * (-10)))
}
