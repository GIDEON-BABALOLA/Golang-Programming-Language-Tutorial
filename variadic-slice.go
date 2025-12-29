package main

import "fmt"

// VARIADIC
// Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.
// A variadic function receives the variadic arguments as a slice
func sumVariadic(nums ...int) int {
	// nums is just a slice
	num := 0
	for i := 0; i < len(nums); i++ {
		num = num + nums[i]
	}
	return num
}
// SPREAD OPERATOR
// The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.
func printStrings(strings ...string){
for i := 0; i < len(strings); i++ {
	fmt.Println(strings[i])
}
}
func sumOfCosts(nums ...float64) float64 {
total := 0.0
for i := 0; i < len(nums); i++ {
	total = total + nums[i]
}
return total
}
func testOfCosts(nums ...float64) {
	total := sumOfCosts(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("======== END REPORT =======")
}
func VariadicSlice(){
newSliced := []int{1, 2, 3}
total := sumVariadic(1, 2, 3)
newTotal := sumVariadic(newSliced...) // spread operator
names := []string{"bob", "sue", "alice"} // spread operator
fmt.Println(total)
fmt.Println(newTotal)
// prints "6"
printStrings(names...)
testOfCosts(1.0, 2.0, 3.0)
}