package main

import (
	"errors"
	"fmt"
)

// SLICES IN GO
// 99 times out of 100 you will use a slice instead of an array when working with ordered lists.
// Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.
// A slice is a dynamically-sized, flexible view of the elements of an array
// Slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:
const ( 
planFree = "free" 
planPro = "pro"
) 
func getMessageWithRetriesForPlan(plan string) ([]string, error) {
allMessages := getMessageWithRetries();
if plan == planPro {
	return allMessages[:], nil
}
if plan == planFree {
	return allMessages[0:2], nil
}
return nil, errors.New("unsupported plan")
}
func getMessageWithRetriesSlice() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}
func testSlices(name string, doneAt int, plan string) {
	defer fmt.Println("========================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()
	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	for i := 0;  i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded")
		}
		if i == len(messages) - 1 {
			fmt.Println("no response")
		}
	}
}
func SlicesInGo(){
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
fmt.Println(mySlice)
fmt.Println(primes[:])
fmt.Println(getMessageWithRetriesForPlan(planFree))
testSlices("Ozgur", 3, planFree)
testSlices("Jeff", 3, planPro)
testSlices("Sally", 2, planPro)
testSlices("Sally", 3, "no plan")

// mySlice = {3, 5, 7}
}