// Go forces us to think of each individual error that is passed back from a dangerous function
package main

import "fmt"

// Errors in Go are just values, they are just interfaces
// Go programs express errors with error values. An Error is any tyoe that implement the simple built in error interface
// The built in error interface is an interface with a simgle method, which is the built in error method
// When you return a non-nil error in Go, it's convention to return the "zero" values of all other return values
// Being good at formatting errors in Go is the same as being good at formatting errors in strings, since an error is just a string or nil %v is the default formatter
type error interface {
	Error() string
}
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error){
costForCustomer, err := sendSMS(msgToCustomer)
if err != nil {
	return 0.0, err
}
costForSpouse, err := sendSMS(msgToSpouse)
if err != nil {
	return 0.0, err
}
return costForCustomer + costForSpouse, nil
}
func sendSMS(message string) (float64, error){
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}
func testError(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("==============")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}
func ErrorHandling(){
	testError("Thanks for coming in to our flower shop", "We hope you enjoyed your gift.")
	testError("Thanks for joining us!", "Have a good day.")
}