package main

import "fmt"

// Base interface
type Device interface {
	TurnOn()
	TurnOff()
}

// Another base interface
type Speaker interface {
	PlaySound()
}

// Embedded interface that inherits from Device and Speaker
type SmartDevice interface {
	Device
	Speaker
	ConnectToWiFi()
}

// Concrete struct implementing the embedded interface
type SmartSpeaker struct {
	Name string
}

// Implementing Device methods
func (s SmartSpeaker) TurnOn() {
	fmt.Println(s.Name, "is now ON")
}

func (s SmartSpeaker) TurnOff() {
	fmt.Println(s.Name, "is now OFF")
}

// Implementing Speaker method
func (s SmartSpeaker) PlaySound() {
	fmt.Println(s.Name, "is playing music")
}

// Implementing SmartDevice method
func (s SmartSpeaker) ConnectToWiFi() {
	fmt.Println(s.Name, "connected to WiFi")
}

func main() {
	var mySpeaker SmartDevice = SmartSpeaker{Name: "Echo Dot"}

	mySpeaker.TurnOn()
	mySpeaker.ConnectToWiFi()
	mySpeaker.PlaySound()
	mySpeaker.TurnOff()
}
