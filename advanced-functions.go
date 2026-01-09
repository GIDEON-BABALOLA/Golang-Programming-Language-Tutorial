package main
import "fmt"
// DEFINITION: FIRST-CLASS FUNCTIONS
// A first-class function is a function that can be treated like any other value. Go supports first-class functions. A function's type is dependent on the types of its parameters and return values. For example, these are different function types:
// func() int
// func(string) int

// DEFINITION: HIGHER-ORDER FUNCTIONS
// A higher-order function is a function that takes a function as an argument or returns a function asa return value. Go supports higher-order functions. For example, this function takes a function as an argument.
// func aggregate(a, b, c, int, arithmetic func(int, int) int) int
func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}
func addSignature(message string) string {
	return message + "Kind regards."
}
func addGreeting(message string) string {
	return "Hello!" + message
}
func testAdvancedFunctions(messages []string, formatter func(string) string){
	defer fmt.Println("=============================================")
}
func AdvancedFunctions(){

}