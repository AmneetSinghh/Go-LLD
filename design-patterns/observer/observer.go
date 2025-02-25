package main

import "fmt"

/*

The Observer Pattern is used when multiple objects (observers) need to be notified of changes in another object (subject) without tightly coupling them.

✅ When to Use?

When an object's state changes, and multiple dependent objects need to be updated automatically.
When you want to decouple the subject (publisher) and observers (subscribers).
When you have a one-to-many relationship.

ex: kafka, GUi button handlers etc.

*/

type Observer interface {
	update(string)
	getId() string
}

type Producer interface {
	register(Observer)
	deregister(Observer)
	notifyAll()
}

type Kafka struct {
	oberverList []Observer
	name        string
	inStock     bool
}

func newKafka(name string) *Kafka {
	return &Kafka{
		name: name,
	}
}
func (i *Kafka) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *Kafka) notifyAll() {
	for _, observer := range i.oberverList {
		observer.update(i.name)
	}
}
func (i *Kafka) register(o Observer) {
	i.oberverList = append(i.oberverList, o)
}
func (i *Kafka) deregister(o Observer) {
	i.oberverList = removeFromSlice(i.oberverList, o)
}

// utils
func removeFromSlice(observerList []Observer, observerRm Observer) []Observer {
	l := len(observerList)
	for i, ooo := range observerList {
		if observerRm.getId() == ooo.getId() {
			observerList[i] = observerList[l-1] // lastremoves here. remove last.
			return observerList[:l-1]
		}
	}
	return observerList
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getId() string {
	return c.id
}

func main() {

	shirtItem := newKafka("Nike Shirt")

	var observerFirst Observer = &customer{id: "abc@gmail.com"}
	var observerSecond Observer = &customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
