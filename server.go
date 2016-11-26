// server.go
package main

import (
	"fmt"
)

type Stereotype int

const (
	TypicalNoob Stereotype = iota
	SecondOne
	ThirdOne
)

func main() {
	fmt.Println("Hello World!", TypicalNoob)
	fmt.Println("Hello World!", SecondOne)
	fmt.Println("Hello World!", ThirdOne)
	const xy = iota
	const ab = iota
	var qw = iota
	fmt.Println("Hello World!", xy)
	fmt.Println("Hello World!", ab)

}
