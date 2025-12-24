package main
import "fmt"
func getLength(s string) int {
	return len(s)
}
func conditionals(){
	// Here are the comparison operators in Go
	// == equal to
	// != not equal to
	// < less than
	// > greater than
	// <= less than or equal to
	// >= greater than or equal to
	messageLen := 10
	maxMessageLen := 20 
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen);
	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	}else {
		fmt.Println("Message not sent")
	}
   // We can do this in Golang, now length variable can only be used within the if statement 
   email := ""
   if length := getLength(email); length < 1 {
	fmt.Println("Email is invalid")
   }
}