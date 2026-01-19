package main
import "fmt"
// ANONYMOUS FUNCTIONS
// Anonymous functions are true to form in that they have no name.We've been using them throughout this chapter, but we haven't really talked about them yet.
// Anonymous functions are useful when defining a function that will only be used once or to create a quick closure

// doMath accepts a function that converts one int to another
// and a slice of ints. It returns a slice of ints that have been converted by the passed in function
func doMath (f func(int) int, numbers []int) []int {
	result := []int{};
	for _, num := range numbers {
		result = append(result, f(num))
	}
	return result;
}
// Complete the printReports functin
func printReports(messages []string) {
for _, msg := range messages {
printCostReport(func(text string) int {
  return 2 * len(msg)
	}, msg)
}
}
// don't touch below this line
func testAnonymous(messages []string){
defer fmt.Println("============================")
printReports(messages)
}
func printCostReport(costCalculator func(string) int, message string){
cost := costCalculator(message)
fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
fmt.Println()
}
func AnonymousFunctions(){
testAnonymous([]string{
	"Here's Johnny!",
	"Go ahead, make my day",
	"You had me at hello",
	"There's no place like home",
})
testAnonymous([]string{
	"My name is Gideon babalola",
	"May the Force be with you",
	"Show me the money",
	"Go ahead, make m y day",
})
nums := []int{1, 2, 3, 4, 5}
// Here we define an anonymous function that doubles an int and passit to doMath
allNumbersDoubled := doMath(func(x int) int {
	return x + x
}, nums)
fmt.Println(allNumbersDoubled)
//  
}
// https://chatgpt.com/share/6966ecbf-d6f0-8013-b387-6115ad3e7b26
// dont fuck with me