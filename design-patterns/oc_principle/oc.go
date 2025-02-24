package main

import "fmt"

// whenever we need to add new case, we need to change existing code.
// violation of open closed principle.
func calculateShippingCost(carrier string, weight int) int {
	switch carrier {
	case "UPS":
		return weight * 1
	case "FedEx":
		return weight * 2
	case "DHL":
		return weight * 3
	default:
		return weight * 5
	}
}

// **** OC ******
// make code adhere to open closed principle

type ShippingStrategy interface {
	calculate(weight int) int
}

type UPS struct{}

func (u *UPS) calculate(weight int) int {
	return weight * 1
}

type Fedx struct{}

func (u *Fedx) calculate(weight int) int {
	return weight * 2
}

type DHL struct{}

func (u *DHL) calculate(weight int) int {
	return weight * 3
}

func calculateShipingCost(carriers []ShippingStrategy) {
	for _, carrier := range carriers {
		fmt.Println(carrier, carrier.calculate(12))
	}
}

func main() {
	calculateShipingCost([]ShippingStrategy{&UPS{}, &Fedx{}, &DHL{}})
}
