// Arrays have a fixed size in Go
package main

import "fmt"

// Arrays are fixed-size groups of variables of the same type
// The type [n]T is an array of n values of the type T
// To declare an array of 10 integers
// To declare an array of 10 integers:
var myInts [10]int
func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}
func sendArray(name string, doneAt int){
	fmt.Printf("sending to %v...", name)
	fmt.Println()
	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
        msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded")
			break
		}
		if i == len(messages) - 1 {
			fmt.Println("complete failure")
		}
	}
} 
func ArraysInGo(){
// or to declare an initialized literal
// primes := [6]int{2, 3, 7, 11, 13}
sendArray("Bob", 0)
sendArray("Alice", 1)
sendArray("Mangalam", 2)
sendArray("Ozgur", 3)
}