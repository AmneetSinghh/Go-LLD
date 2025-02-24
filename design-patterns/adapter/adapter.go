package main

import "fmt"

type computer interface {
	insertIntoLighteningPort()
}

type oldComputer interface {
	insertIntoUSBPort()
}

type Mac struct{}

func (m *Mac) insertIntoLighteningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type window struct{}

// implementing old computer because its working is that.
// lets say this is another team code, you can change te function name or something....

func (m *window) insertIntoUSBPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type windowAdapter struct {
	windowMachine oldComputer
}

func (wc *windowAdapter) insertIntoLighteningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	wc.windowMachine.insertIntoUSBPort()
}

type client struct{}

func (c *client) insertLighteningConnector(com computer) {
	com.insertIntoLighteningPort()
}

func main() {
	c := &client{}
	mac := &Mac{}

	c.insertLighteningConnector(mac)

	windowAdpter := &windowAdapter{
		windowMachine: &window{}, // concrete implementation
	}
	c.insertLighteningConnector(windowAdpter)
}

/*

- A structural design pattern that converts one interface into another that a client expects, allowing incompatible interfaces to work together.
- Software Integration: Wrapping legacy systems (e.g., logging or payment processing) to match modern interfaces.

- If a third-party interface is fixed and cannot be changed because it is legacy, but our new interface is different
and our clients expect to use our interface, we create an adapter. This adapter internally calls the legacy APIs while exposing an
interface that conforms to our new architecture.


*/
