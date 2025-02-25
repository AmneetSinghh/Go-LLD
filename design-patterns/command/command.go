package main

import "fmt"

/*
Decoupling: The remote (invoker) holds command objects rather than directly interacting with the TV (receiver).
Encapsulation: Each command encapsulates a specific TV operation (e.g., on, off, volume adjustments).
Flexibility: Commands can be dynamically assigned or swapped, allowing the same remote to control different devices. ( TV or DVD)
Undo Capability: Commands can store state for supporting undo/redo actions.
Loose Coupling: The remote doesn't need to know the implementation details of the TV, ensuring high modularity. ( DI principle )


*/

/*

	// command will interact with product not concrete implementin of product
	// remote contorl interact with comand, lets say onCommand. but internaly it can be TV of in its implementation.
	// commands are the wrappers, which wraps TV internal operations......

	RemoteControl 				-> 			Commands 				-> Product
		|
	concrete implementations			concrete imple..				concrete impllementions///
		GameControl							onCommand						actuall on
		DvdControl							offCommand						actuall off
		TvControl							upCommand						actual up
											downCommand						actual down


*/

type Command interface {
	Execute()
	Undo()
}

// can be DVD
type Product interface {
	TurnOn()
	TurnOff()
	VolumeUp()
	VolumeDown()
}

// product implements.
type Television struct {
	on     bool
	volume int
}

func newTelevision() *Television {
	return &Television{
		on:     false,
		volume: 10,
	}
}

func (tv *Television) TurnOn() {
	tv.on = true
	fmt.Println("Tv is on")
}
func (tv *Television) TurnOff() {
	tv.on = false
	fmt.Println("Tv is off")
}
func (tv *Television) VolumeUp() {
	tv.volume++
	fmt.Println("volume up")
}
func (tv *Television) VolumeDown() {
	tv.volume--
	fmt.Println("volume down")
}

// list of commands

// on command
type TurnOnCommand struct {
	product Product
}

func newTurnOnCommand(p Product) *TurnOnCommand {
	return &TurnOnCommand{product: p}
}
func (on *TurnOnCommand) Execute() {
	on.product.TurnOn()
}
func (on *TurnOnCommand) Undo() {
	on.product.TurnOff()
}

// off command
type TurnOffCommand struct {
	product Product
}

func newTurnOffCommand(p Product) *TurnOffCommand {
	return &TurnOffCommand{product: p}
}
func (on *TurnOffCommand) Execute() {
	on.product.TurnOff()
}
func (on *TurnOffCommand) Undo() {
	on.product.TurnOn()
}

// up command
type VolumeUpCommand struct {
	product Product
}

func newVolumeUpCommand(p Product) *VolumeUpCommand {
	return &VolumeUpCommand{product: p}
}
func (on *VolumeUpCommand) Execute() {
	on.product.VolumeUp()
}
func (on *VolumeUpCommand) Undo() {
	on.product.VolumeDown()
}

// volun down command
type VolumedowCommand struct {
	product Product
}

func newVolumedowCommand(p Product) *VolumedowCommand {
	return &VolumedowCommand{product: p}
}
func (on *VolumedowCommand) Execute() {
	on.product.VolumeDown()
}
func (on *VolumedowCommand) Undo() {
	on.product.VolumeUp()
}

type Controller interface {
	pressOnButton()
	pressOffButton()
	pressUpButton()
	pressDownButton()
	PressUndoButton()
	setCommand(Command, Command, Command, Command)
}

// interface controller is needed, otherwise outside functions, cn call oncommand.execute by their own way.....

type RemoteControl struct {
	onCommand         Command
	offCommand        Command
	volumeUpCommand   Command
	volumeDownCommand Command
	commandHistory    []Command
}

func NewRemoteControl() Controller {
	var controller Controller = &RemoteControl{
		commandHistory: []Command{},
	}
	return controller
}

func (r *RemoteControl) setCommand(on Command, off Command, vup Command, vdown Command) {
	r.onCommand = on
	r.offCommand = off
	r.volumeDownCommand = vdown
	r.volumeUpCommand = vup
}
func (r *RemoteControl) pressOnButton() {
	r.onCommand.Execute()
	r.commandHistory = append(r.commandHistory, r.onCommand)
}
func (r *RemoteControl) pressOffButton() {
	r.offCommand.Execute()
	r.commandHistory = append(r.commandHistory, r.offCommand)
}
func (r *RemoteControl) pressUpButton() {
	r.volumeUpCommand.Execute()
	r.commandHistory = append(r.commandHistory, r.volumeUpCommand)
}
func (r *RemoteControl) pressDownButton() {
	r.volumeDownCommand.Execute()
	r.commandHistory = append(r.commandHistory, r.volumeDownCommand)
}
func (r *RemoteControl) PressUndoButton() {
	if len(r.commandHistory) == 0 {
		fmt.Println("No command to undo")
		return
	}
	lastIndex := len(r.commandHistory) - 1
	lastCommand := r.commandHistory[lastIndex]
	r.commandHistory = r.commandHistory[:lastIndex]
	fmt.Println("undowing last command..")
	lastCommand.Undo()
}

func main() {
	// commands
	var p Product = &Television{}
	on := newTurnOnCommand(p)
	off := newTurnOffCommand(p)
	vDown := newVolumeUpCommand(p)
	vUp := newVolumedowCommand(p)
	r := NewRemoteControl()
	r.setCommand(on, off, vUp, vDown)

	r.pressOnButton()
	r.pressOffButton()
	r.pressUpButton()
	r.pressDownButton()

	r.PressUndoButton()
}
