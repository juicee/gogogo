package main

import (
	"fmt"
)

func main() {
	z := 37
	pi := &z
	ppi := &pi

	fmt.Println(z, pi, *ppi)

	mt := map[int]int{}
	mt[1] = 2
	mt[3] = 4
	fmt.Println(mt)

	list := []string{"123", "abc", "kdo"}

	for value := range mt {
		fmt.Printf(" -> %d\n", value)
	}

	for value := range list {
		fmt.Printf(" -> %s\n", value)

	}
}
