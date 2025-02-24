package main

import "fmt"

/*
- The Builder Pattern ensures that the object is fully constructed before it's used, preventing an inconsistent state.
- In build() step, I can write validation logic ( to check whether all fields are set or not )
- its not looking clutered.
- when we initilize construct with multiple parameters, its look clutered, and doesn't follow SOLID principles
- Single line, improves readibility


// First method :

client responsibilty to make method chaining, so don't forgot to miss some set method. its the responsibilty of teh person
*/

type House struct {
	window int
	door   int
	color  int
}

type HouseBuilder struct {
	house House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (hb *HouseBuilder) setWindow(w int) *HouseBuilder {
	hb.house.window = w
	return hb
}

func (hb *HouseBuilder) setDoor(d int) *HouseBuilder {
	hb.house.door = d
	return hb
}

func (hb *HouseBuilder) setColor(c int) *HouseBuilder {
	hb.house.color = c
	return hb
}

func (hb *HouseBuilder) Build() *House {
	// can make validations in this.
	if hb.house.window == 0 || hb.house.door == 1 {
		return nil
	}
	return &hb.house
}

func main_1() {
	house := NewHouseBuilder().
		setColor(12).
		setDoor(12).
		setWindow(1).
		Build()
	fmt.Println(house)
}

/*
Director method :

- problem :

In above client is making builder object, but he can do mistake as well. So lets say if we have only 3 types of builders universal,
other builders are not possible, then just give responseibility of builder to director. So director will make builder objects.
we just need to give type of builder to director object.

- ✅ Director ensures construction correctness.
- when in multilpe places we are using builders, then director might be helpful. not in every case.


*/

type IBuilder interface {
	setWindow()
	setDoor()
	setColor()
	getHouse() House
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoor()
	d.builder.setWindow()
	d.builder.setColor()
	return d.builder.getHouse()
}

type Builder_1 struct {
	window int
	door   int
	color  int
}

func newBuilder_1() *Builder_1 {
	return &Builder_1{}
}
func (b *Builder_1) setWindow() {
	b.window = 1
}
func (b *Builder_1) setDoor() {
	b.door = 2
}
func (b *Builder_1) setColor() {
	b.color = 3
}
func (b *Builder_1) getHouse() House {
	return House{
		door:   b.door,
		window: b.window,
		color:  b.color,
	}
}

type Builder_2 struct {
	window int
	door   int
	color  int
}

func newBuilder_2() *Builder_1 {
	return &Builder_1{}
}
func (b *Builder_2) setWindow() {
	b.window = 12
}
func (b *Builder_2) setDoor() {
	b.door = 13
}
func (b *Builder_2) setColor() {
	b.color = 14
}
func (b *Builder_2) getHouse() House {
	return House{
		door:   b.door,
		window: b.window,
		color:  b.color,
	}
}

type BuilderType string

const (
	Builder1 BuilderType = "builder_1"
	Builder2 BuilderType = "builder_2"
)

// factory.
func getBuilder(bt BuilderType) IBuilder {
	switch bt {
	case Builder1:
		return newBuilder_1()
	case Builder2:
		return newBuilder_2()
	default:
		return nil
	}
}

func main() {
	builder1 := getBuilder(Builder1)
	builder2 := getBuilder(Builder2)

	director := newDirector(builder1)
	house := director.buildHouse()

	fmt.Printf("Normal House Door Type: %d\n", house.door)
	fmt.Printf("Normal House Window Type: %d\n", house.window)
	fmt.Printf("Normal House Num Floor: %d\n", house.color)

	director.setBuilder(builder2)
	iglooHouse := director.buildHouse()
	fmt.Printf("iglooHouse Door Type: %d\n", iglooHouse.door)
	fmt.Printf("iglooHouse Window Type: %d\n", iglooHouse.window)
	fmt.Printf("iglooHouse Num Floor: %d\n", iglooHouse.color)

}
