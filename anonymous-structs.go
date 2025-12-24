package main
import "fmt"
// An anonymous struct is just like a normal struct but it is defined without a name and therefore cannot be referenced elsewhere in the code
// To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after decalring the type
// The only reason you use anonymous structs is when you have no reason to create more than one instance of the struct. 
// You can also nest anonymous structs as fields within other structs
type car struct {
	Make string
	Model string
	Height int
	Width int
	// Wheel is a filed containing an anonymous struct 
	Wheel struct {
		Radius int
		Material string
	}
}
func AnonymousStructs(){
myCar := struct {
Make string
Model string
} {
	Make: "tesla",
	Model: "model 3",
}
fmt.Println(myCar)
var myGoodCar car = car{}
myGoodCar.Height = 5
fmt.Println(myGoodCar)
}