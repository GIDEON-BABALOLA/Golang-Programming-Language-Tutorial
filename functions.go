package main
import "fmt"
func concat(s1 string, s2 string) string {
	return s1 + s2
}
// Multiple parameters, when multiple arguments are of the same type, the type only needs to be declared after the last one, assuming they are in order
func add(x, y int) int {
	return x + y
}
func incrementSends (sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}
// Functions in Go can have multiple return values, and when this happens we specify the type of return values in parentheses
func getNames() (string, string){
return "John", "Doe"
}
func getPoint() (x int, y int)  {
	return 3, 4
}
// In Golang we can name our return values, and if we do, it alters the behavior of the function just a little bit
func getCoords() (x, y int){
	// x and y are initialized with their zero values
	return // automatically returns x and y
} 
// Is the same as
func getCoordsAgain() (int, int) {
	var x int
	var y int
	return x,  y
}
// Boot.dev Recommendations concerning named return values, he recommends using named returns when you want to document what the intended purpose of the return value is
func yearsUntilEvents(age int) ( yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age;
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age;
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	// return
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}
func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====================================================")
}
// In Golang variables are passed by value, not by reference
func functions(){
fmt.Println(concat("Lane,", " happy birthday"))
fmt.Println(concat("Elon,", " hope that tesla thing works out"))
fmt.Println(concat("Go, ", "is fantastic"))
fmt.Println(add(1, 2))
sendsSoFar := 430
const sendsToAdd = 25
sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
fmt.Println("you've sent", sendsSoFar, "messages")
firstname, _ := getNames()
fmt.Println("Welcome to Textio,", firstname)
// Functions in Go can have multiple return values, Golang does not allow you to have unused variables, and also Go allows a function to return multiple values and there can be situations whereby you will only need a value from it, the multiple values that is returned by Golang
x, _ := getPoint();
fmt.Println(x)
coordinates, _ := getCoords()
coordinatesAgain, _ := getCoordsAgain()
fmt.Println(coordinates, coordinatesAgain)
test(4)
test(10)
test(22)
test(35)
}