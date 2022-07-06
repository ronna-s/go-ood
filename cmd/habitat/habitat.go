package main

import (
	"fmt"

	"github.com/ronnas/go-ood/pkg/animal"
)

func main() {
	h := Habitat{Name: "The animal coworking space"}
	fmt.Println(h.Add(animal.NewDeer()))
	fmt.Println(h.Add(animal.NewDeer()))
	fmt.Println(h.Add(animal.NewLion()))
	fmt.Println(h.Add(animal.NewDeer()))
}

// Habitat represents the habitat of a group of animals
type Habitat struct {
	Animals []animal.Animal
	Name    string
}

// Add adds an animal to the habitat
func (h *Habitat) Add(a animal.Animal) string {
	for i, a2 := range h.Animals {
		if p, ok := a.(animal.Predator); ok {
			if p.Eats(a2) && !p.Full() {
				p.Eat(a2)
				h.Animals = append(h.Animals[:i], h.Animals[i+1:]...)
				h.Animals = append(h.Animals, a)
				return p.MakeSound()
			}
		}
		if p, ok := a2.(animal.Predator); ok {
			if p.Eats(a) && !p.Full() {
				p.Eat(a)
				return a.MakeSound()
			}
		}
	}
	h.Animals = append(h.Animals, a)
	return a.MakeSound()
}
