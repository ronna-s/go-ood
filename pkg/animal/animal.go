package animal

import "time"

// Now ...
var Now = time.Now

type Animal interface {
	MakeSound() string
}

type Predator interface {
	Animal
	Eats(Animal) bool
	Eat(Animal) string
	Full() bool
}
type BaseAnimal struct {
	Sound string
}

func (a BaseAnimal) MakeSound() string {
	return a.Sound
}

type Deer struct {
	BaseAnimal
}
type Lion struct {
	BaseAnimal
	LastFed time.Time
}

func (l Lion) Eats(a Animal) bool {
	if _, eats := a.(Deer); eats {
		return true
	}
	return false
}
func (l *Lion) Eat(_ Animal) string {
	l.LastFed = Now()
	return l.MakeSound()
}
func (l Lion) Full() bool {
	return l.LastFed.After(Now().Add(-1 * time.Hour))
}

func NewLion() *Lion {
	return &Lion{BaseAnimal: BaseAnimal{Sound: "Roar!"}}
}
func NewDeer() Deer {
	return Deer{BaseAnimal: BaseAnimal{Sound: "I'm so cute!"}}
}
