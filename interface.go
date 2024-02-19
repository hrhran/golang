package main

import "fmt"

type engine interface {
	distanceLeft() float32
}

type gasEngine struct {
	kmpl   uint8
	petrol uint8
}

type electricEngine struct {
	kpkwh uint8
	kwh   uint8
}

func (g gasEngine) distanceLeft() float32 {
	return float32(g.kmpl / g.petrol)
}

func (e electricEngine) distanceLeft() float32 {
	return float32(e.kpkwh / e.kwh)
}

func canMakeIt(e engine, kilometers uint8) {
	if e.distanceLeft() > float32(kilometers) {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't make it!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{20, 3}
	canMakeIt(myEngine, 100)
}
