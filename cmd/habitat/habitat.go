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
}
