package main
import "fmt"
func formatting() {
	// Using %v - general use for any data types
	fmt.Printf("I am %v years old\n", 10)
	fmt.Printf("I am %v years old\n", "way too many")
	// Using %s - this interpolates a string
	fmt.Printf("I am %s years old\n", "way too many")
	// Using %d - this interpolates a integer in decimal form
	fmt.Printf("I am %v years old\n", 10)
	// Using %f - this is for floats so it interpolates a decimal
	fmt.Printf("I am %f years old\n", 10.523)
	// Using %f - The ".2" rounds the number to 2 decimal places
	fmt.Printf("I ma %.2f years old\n", 10.526)
	// Example
	const name = "Saul Goodman"
	const openRate = 30.5
	// msg := "Hi %s, your open rate is %.1f percent"
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)
	// fmt.Println(msg, name, openRate)
	fmt.Println(msg)
}