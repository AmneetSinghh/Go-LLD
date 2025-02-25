package main

import (
	"fmt"
	"log"
)

/*

✅ Use the State Pattern When:
- An object’s behavior depends on its state and changes dynamically.
- You want to avoid complex if-else or switch statements for handling different states.
- The system has clearly defined states with distinct behaviors (e.g., Vending Machine, Traffic Light, Order Processing).
- You need to add new states without modifying existing code (Open/Closed Principle - OCP).
- Each state should be encapsulated separately for better readability & maintainability.
- Transitions between states should be explicit and managed dynamically.

🚫 Avoid the State Pattern When:
- The number of states is fixed and unlikely to change.
- A simple if-else condition is enough (e.g., 2-3 states with minimal logic).
- Performance is a top priority, and extra abstraction causes overhead.

*/

/*
- first use simple if else.
- question yourself, can this if else increase in future? if it can icrease its violatoin of os principle, we need state pattern.
- now in state pattern , quetion yourself can these, number of states can increase to 100s... if yes, we need to segregate and make smaller interfaces
// so it must adher to lisbok and interface seg.. princpiple.
*/
type VendingMachine struct {
	hasItem            State
	itemRequestedState State
	hasMoney           State
	noItem             State
	currentState       State
	itemCount          int
	itemPrice          int
}

func newVendingMachine(itemCount, itemPrice int) *VendingMachine {

	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	noItemState := &NoItemState{VendingMachine: v}
	hasItemState := &HasItemState{VendingMachine: v}
	itemReqState := &ItemRequestedState{VendingMachine: v}
	hasMoneyState := &HasMoneyState{VendingMachine: v}
	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequestedState = itemReqState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}
func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}
func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}
func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}
func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type NoItemState struct {
	VendingMachine *VendingMachine
}

func (i *NoItemState) addItem(count int) error {
	i.VendingMachine.incrementItemCount(count)
	i.VendingMachine.setState(i.VendingMachine.hasItem)
	return nil
}
func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item ut of stock")
}
func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

type HasItemState struct {
	VendingMachine *VendingMachine
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.VendingMachine.incrementItemCount(count)
	return nil
}
func (i *HasItemState) requestItem() error {
	if i.VendingMachine.itemCount == 0 {
		i.VendingMachine.setState(i.VendingMachine.noItem)
		return fmt.Errorf("item out of stock")
	}
	fmt.Println("Items requested")
	i.VendingMachine.setState(i.VendingMachine.itemRequestedState)
	return nil
}
func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}
func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item alreasdy processed")
}
func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.VendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.VendingMachine.itemPrice)
	}
	fmt.Println("money entereed is OK")
	i.VendingMachine.setState(i.VendingMachine.hasMoney)
	return nil
}
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}
func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}
func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("ite out of stock")
}
func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing item")
	i.VendingMachine.itemCount = i.VendingMachine.itemCount - 1
	if i.VendingMachine.itemCount == 0 {
		i.VendingMachine.setState(i.VendingMachine.noItem)
	} else {
		i.VendingMachine.setState(i.VendingMachine.hasItem)
	}
	return nil
}

func main() {
	vendingMachine := newVendingMachine(1, 10)
	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(100)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
