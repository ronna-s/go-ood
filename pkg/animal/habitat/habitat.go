package habitat

import "github.com/ronnas/go-ood/pkg/animal"

type Habitat struct {
	Animals []animal.Animal
	Name    string
}

func New(name string) Habitat {

}

func (h *Habitat) Add(a animal.Animal) (res []string) {
	for i := range h.Animals {
		if h.Animals[i].Eats(a) {
			res = append(res, h.Animals[i].MakeSound())
		} else if a.Eats(h.Animals[i]) {
			res = append(res, a.MakeSound())
			h.Animals = append(h.Animals[:i], h.Animals[:i+1])
			i--
		}
	}
	if len(res) == 0 {
		res = []string{a.MakeSound()}
	}
	return res
}
