package main

import "fmt"

type Gopher struct {
	name          string
	age           int
	doritos_eaten int
	capacity      int
}

func (g *Gopher) eat_dorito() {
	g.doritos_eaten++
}

// TODO: Implement is_full() which will return true if your gopher is full!

func main() {
	// TODO make a gopher!
	g := Gopher{}

	for {
		g.eat_dorito()
		if g.is_full() {
			break
		}
	}

	fmt.Printf("%s ate %d doritos. He's full!", g.name, g.doritos_eaten)
}

// TO DO Bonus!
// Reformat the code to have an EMBEDDED STRUCT.
/*
type Animal struct {
	name          string
	age           int
}
type Gopher struct {
	Animal
	doritos_eaten int
	capacity      int
}
*/

// The types within type Animal will be "promoted"
// to the type Gopher. i.e. after initializing properly
// (HINT: https://travix.io/type-embedding-in-go-ba40dd4264df)
// you will not need to do "g.Animal.name", you can just use "g.name"!
// Cool!

