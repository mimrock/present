package main

import (
	"fmt"
)

type Truck struct {
	Cargo float64
}

func (t *Truck) Report() {
	fmt.Println("My cargo weights", t.Cargo, "kg")
}

func main() {
	goliath := Truck{75.5}
	goliath.Report()
}