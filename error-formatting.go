package main
// Being good at formatting errors in Go is the same as being good at formatting errors in strings, since an error is just a string or nil %v is the default formatter
import "fmt"
func getSMSErrorString(cost float64, reciepient string) string{
return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, reciepient)
}
func testErrorFormatting(cost float64, reciepient string){
s := getSMSErrorString(cost, reciepient)
fmt.Println(s)
fmt.Println("===============================")
}
func ErrorFormatting(){
testErrorFormatting(1.4, "+1 (435) 555 0923")
testErrorFormatting(2.1, "+2 (702) 555 3452")
testErrorFormatting(32.1, "+1 (801) 555 7456")
}