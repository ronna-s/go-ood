package main

import (
	"fmt"

	"github.com/ronnas/go-ood/pkg/sync"
)

func main() {
	m := sync.Map[string, int]{M: map[string]int{"hello": 1}}
	fmt.Println(m.At("hello"))
	//fmt.Println("hello world!", animal.Lion{})
}
