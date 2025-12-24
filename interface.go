// An interface in Go is just a collection of method signatures, not data types
// Interface are collection of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.
// In the following example, a "shape" must be able to return its area and perimeter. Both rect and circle fulfill the interface
// Multiple types can implement the same interface
package main

import (
	"fmt"
	"math"
	"time"
)

// import "fmt"
type shape interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	width, height float64
}
type message interface {
getMessage() string
}
type employee interface {
	getName() string
	getSalary() int
}
type contractor struct {
	name string
	hourlyPay int
	hoursPerYear int
}
type fullTime struct {
	name string
	salary int
}
func sendMessage(msg message){
fmt.Println(msg.getMessage())
}
type birthdayMessage struct {
	birthdayTime time.Time
	recipientName string
}
type sendingReport struct {
	reportName string
	numberOfSends int
}
func  (r rectangle) area() float64 {
return r.width * r.height
}
func (r rectangle) perimeter() float64 {
return 2*r.width + 2*r.height
}
type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}
func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v`, sr.reportName, sr.numberOfSends)
}
func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}
func (ft fullTime) getSalary() int {
return ft.salary
}
func (ft fullTime) getName() string {
	return ft.name
}
func testAgain(m message) {
	sendMessage(m)
	fmt.Println("====================")
}
func testNow(e employee){
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("===================================")
}
// Both the rectangle and the circle implements the shape interface
func Interface() {
// So the good thing about the interface now is that testAgain accepts message interface as its parameter and the only reason why sendingReport can be passed as an argument is because sendingReport satisfies the interface getMessage() so if getMessage() is removed from sendingReport then it means that testAgain cannot accept an argument such as sendingReport
testAgain(sendingReport{
	reportName: "First Report",
	numberOfSends: 10,
})
testAgain(birthdayMessage{
	recipientName: "John Doe",
	birthdayTime: time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
})
firstReport := sendingReport{
	reportName: "Full Blood Test",
	numberOfSends: 5,
}
// firstBirthdayMessage := birthdayMessage{

// }
sendMessage(firstReport)
// Here multiple types are implementing the same interface
testNow(fullTime{
	name: "David",
	salary: 50000,
})
testNow(contractor{
	name: "Bob", 
	hourlyPay: 100,
	hoursPerYear: 73,
})
testNow(contractor{
	name: "Jill",
	hourlyPay: 872,
	hoursPerYear: 982,
})
}