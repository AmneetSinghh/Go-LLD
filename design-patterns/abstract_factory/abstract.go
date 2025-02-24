package main

import "fmt"

/*
- 4 have products
- factory that maintain 4 products.
- object inilization no need switch, one facotry will have information about concrete products.

*/

// product 1.

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(l string) {
	s.logo = l
}
func (s *Shoe) setSize(sz int) {
	s.size = sz
}
func (s *Shoe) getLogo() string {
	return s.logo
}
func (s *Shoe) getSize() int {
	return s.size
}

// concrete products_1
type AdidasShoe struct {
	Shoe
}
type NikeShoe struct {
	Shoe
}

// product 2.

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(l string) {
	s.logo = l
}
func (s *Shirt) setSize(sz int) {
	s.size = sz
}
func (s *Shirt) getLogo() string {
	return s.logo
}
func (s *Shirt) getSize() int {
	return s.size
}

// concrete products_2
type AdidasShirt struct {
	Shirt
}
type NikeShirt struct {
	Shirt
}

// factor will contain both

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

// adidas factory

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{logo: "logo", size: 12},
	}
}
func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{logo: "logo", size: 12},
	}
}

// nike factory

type Nike struct {
}

func (a *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{logo: "logo", size: 12},
	}
}
func (a *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{logo: "logo", size: 12},
	}
}

func getSportFactory(brand string) (ISportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong baed type passed")
}
