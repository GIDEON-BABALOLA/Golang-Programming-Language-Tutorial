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
func AnonymousFunctions(){
nums := []int{1, 2, 3, 4, 5}
// Here we define an anonymous function that doubles an int and passit to doMath
allNumbersDoubled := doMath(func(x int) int {
	return x + x
}, nums)
fmt.Println(allNumbersDoubled)
// prints:
// [2 4 6 8 10]
}