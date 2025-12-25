package main

import "fmt"

// When working with interfaces in Go, every once in a while, you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion
type shapes interface {
area() float64
}
type circles struct {
	radius float64
}
// "c" is a new circle cast from "s"
// which is an instance of a shape.
// "ok" is a bool that is true if s was a circle
// or false if s isn't a circle
func getExpenseReport(e expenses) (string, float64) {
em, ok := e.(emails) // I am type casting e as an interface to type emails
if ok {
	return em.toAddress, em.cost()
}
s, ok := e.(smss) // I am type casting s as an interface to type smss
if ok {
	return s.toPhoneNumber, s.cost()
}
return "", 0.0
}
func (e emails) cost() float64 {
if !e.isSubscribed {
	cost := 0.05 * float64(len(e.body))
	return  cost
}
	cost := 0.01 * float64(len(e.body))
	return  cost
	
}
func (s smss) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}
func (i invalid) cost() float64 {
	return 0.0
}
func myTest(e expenses){
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case emails:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("=================================================")
	case smss:
		fmt.Printf("Report: The smss going to %s will cost: %.2f\n", address, cost)
		fmt.Println("=================================================")
	default:
        fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}
type expenses interface {
	cost() float64
}
type emails struct {
	isSubscribed bool
	body string
	toAddress string
}
type smss struct {
	isSubscribed bool
	body string
	toPhoneNumber string
}
type invalid struct{}
func TypeAssertions(){
// c, ok := s.(circles)
myTest(emails{
	isSubscribed: true,
	body: "hello there",
	toAddress: "john@doe.com",
})
myTest(emails{
	isSubscribed: false,
	body: "This meeting could have been an email",
	toAddress: "jane@doe.com",
})
myTest(emails{
	isSubscribed: false,
	body: "This meeting could have been an email",
	toAddress: "elan@doe.com",
})
myTest(smss{
	isSubscribed: false,
	body: "This meeting could have been an email",
	toPhoneNumber: "+155555509832",
})
myTest(smss{
	isSubscribed: false,
	body: "This meeting could have been an email",
	toPhoneNumber: "+155555504444",
})
myTest(invalid{})
}