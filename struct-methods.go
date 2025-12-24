package main

import "fmt"

// While Go is not object oriented it does suport methods which can be defined on structs. Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function
type rect struct {
	width int
	height int
}
type triangle struct {
	length int
	breath int
}
type authenticationInfo struct {
	username string
	password string
}
// area has a receiver of ( r rect )
func (r rect) area() int {
	return r.height * r.width
} 
// Triangle area has a receiver of  ( t triangle )
func (t triangle) area() int {
return t.length * t.breath
}
func (authI authenticationInfo) getBasicAuth() string {
// Sprintf returns string instead of printing it to standard output or to console
// return "Authorization: Basic " + auth.username + ":" + auth.password
return fmt.Sprintf("Authorization: Basic %s:%s", authI.username, authI.password)
}
func testAuth(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("=============================")
}
func StructMethods() {
r := rect{
	width: 5,
	height: 10,
}
t := triangle{
	length: 5,
	breath: 15,
}
fmt.Println(t.area())
fmt.Println(r.area())
testAuth(authenticationInfo{
	username: "gidiboy",
	password: "gidiboy##",
})
}