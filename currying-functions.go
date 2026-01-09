package main

import (
	"errors"
	"fmt"
)

// CURRYING
// Function currying is the practice of writing a function that takes a function( or functions ) as input, and returns a new function.

// getLogger takes a function that formats two string into
// a single string and returns a function that formats two strings but  prints
// the result instead of returning it.
func getLogger(formatter func(string, string) string) func(string, string){
return func(string1, string2 string) {
	newString := formatter(string1, string2)
	fmt.Println(newString)
}
}
func testCurryingFunction(first string, errors []error, formatter func(string, string) string){
	defer fmt.Println("======================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}
func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}
func CurryingFuncton(){
dbErrors := []error{
	errors.New("out of memory"),
	errors.New("cpu is pegged"),
	errors.New("networking issue"),
	errors.New("invalid syntax"),
}
testCurryingFunction("Error on database server", dbErrors, colonDelimit)
mailErrors := []error{
	errors.New("email too large"),
	errors.New("non alphanumeric symbols found"),
}
testCurryingFunction("Error on mail server", mailErrors, commaDelimit)
squareFunc := selfMath(multiplyF)
doubleFunc := selfMath(addF)
fmt.Println(squareFunc(5))
// prints 25
fmt.Println(doubleFunc(5))
// prints 10
}
func multiplyF(x, y int) int {
	return x * y
}
func addF(x, y int) int {
	return x + y
}
func selfMath(mathFunc func(int, int) int) func (int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}
// In the example above, the selfMath function takes in a function as its paramter, and returns a function that itself returns the value of running that input function on its parameter.