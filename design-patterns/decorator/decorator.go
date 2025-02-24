package main

import "fmt"

type IPizza interface {
	getName() string
}

type VegMania struct{}

func (v *VegMania) getName() string {
	return "veg_mania"
}

// tomato decorator

type TomatoDecorator struct {
	pizza IPizza
}

func (to *TomatoDecorator) getName() string {
	return to.pizza.getName() + "_tomato"
}

// cheese decorator

type CheeseDecorator struct {
	pizza IPizza
}

func (to *CheeseDecorator) getName() string {
	return to.pizza.getName() + "_cheese"
}

func main() {
	pizza := CheeseDecorator{
		pizza: &TomatoDecorator{
			pizza: &VegMania{},
		},
	}
	fmt.Println(pizza.getName())

}
