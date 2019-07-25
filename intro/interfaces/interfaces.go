package main

import "fmt"
import "math"

// TODO: add a method "area" to the interface, and 
// implement it for both circle and rectangles

type geometry interface {
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. 


// Here we implement `geometry` on `rect`s.
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // The `circle` and `rect` struct types both
    // implement the `geometry` interface so we can use
    // instances of
    // these structs as arguments to `measure`.
    measure(r)
    measure(c)
}

