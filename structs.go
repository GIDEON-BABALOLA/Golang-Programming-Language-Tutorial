package main

import (
	"fmt"
)

// The Struct Collection Type
// A struct is a collection type, a collection type is a type that contains other types, a struct is a collection of key value pairs, A struct keys can hold any type not just primitive types like string, integers and booleans
type Car struct {
	Make string
	Model string
	Height int
	Width int
	FrontWheel Wheel
	BackWheel Wheel
}
type Wheel struct {
	Radius int
	Material string
}
type messageToSend struct {
	phoneNumber int
	message string
}
type newMessageToSend struct {
	message string
	sender user
	recipient user
}
type user struct {
	name string
	number int
}
func canSendMessage(mToSend newMessageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	} 
	if mToSend.recipient.name == "" {
		return false
	} 
	if mToSend.sender.number == 0 {
		return false
	} 
	if mToSend.recipient.number == 0 {
		return false
	} 
	return true
}
func tested(m messageToSend){
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("========================================================")
}
func StructsCollectionTypes(){
tested(messageToSend{
	phoneNumber: 148255510981,
	message: "Thanks for signing up",
})
tested(messageToSend{
	phoneNumber: 148255510982,
	message: "Love to have you aboard!",
})
tested(messageToSend{
	phoneNumber: 148255510983,
	message: "We're so exited to have you",
})
fmt.Println(canSendMessage(newMessageToSend{}))
// This is how we can instantiate a new instance of a struct
myCar := Car{}
// when you create myCar as empty {}, all the variables or fields will be initialized as their default value
// The fields of a struct can be accessed using the dot. operator
myCar.FrontWheel.Radius = 5
fmt.Println(myCar)
}