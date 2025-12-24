package main
import "fmt"
type sleakCar struct {
	make string
	model string
}
type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	sleakCar
	bedSize int
}
type userNow struct {
	name string
	number int
}
type sender struct {
	rateLimit int
	userNow
}
func tester(s sender){
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit", s.rateLimit)
	fmt.Println("=============================== ")
}
func EmbeddedStructs(){
lanesTruck := truck{
	bedSize: 10,
	sleakCar: sleakCar{
		make: "toyota",
		model: "camry",
	},
}
fmt.Println(lanesTruck.model, lanesTruck.model, lanesTruck.bedSize)
// Another example
mySender := sender{
	rateLimit: 10,
	userNow: userNow{
		name: "Gideon",
		number: 10,
	},
}
tester(mySender)
tester(sender{
	rateLimit: 15,
	userNow: userNow{
		name: "Jacob",
		number: 13,
	},
})
fmt.Println(mySender)
}