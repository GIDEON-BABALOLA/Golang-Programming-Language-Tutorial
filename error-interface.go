package main

import "fmt"

// Because errors are just interfaces, you can build your own custom types that implement the error interface. Here's an example of a userError struct that implements the error interface
type userError struct {
	name string
}
func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}
// It can then be used here as an error
func canSendToUser(userName string) bool {
	return len(userName) > 10
}
func sendSMSErrorInterface(msg, userName string) (string, error)  {
	if !canSendToUser(userName) {
		return msg, userError{name: userName}
	}
	return msg, nil
}
func testErrorInterface(msg, userName string) {
message, err := sendSMSErrorInterface(msg, userName)
fmt.Println(err.Error()) // Here I can easily access the error() method because it is exposed
if err != nil { 
em, ok := err.(userError)
if ok {
errors := em.Error() 
fmt.Println("Error:", em.name) // But here I have to know the type of error before i can access the properties within it using type asssertions
fmt.Println(errors)
}
return
 }
fmt.Println(message)
}
func ErrorInterface(){
testErrorInterface("Your airtime balance is 50 naira", "Dave")
}