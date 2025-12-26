package main
import "fmt"
// A type switch makes it easy to do several assertions in a series
// A type switch is similar to a regular switch statement, but the cases specify types instead of values
// The reason why we can cast any of the types int, and string to interface num  is because an empty interface in Go is implemented by every types
type person struct {
	name string
}
func printNumbericValue (num interface{}) {
	switch v := num.(type) {
    case int:
    fmt.Printf("%T\n", v)
	case string:
    fmt.Printf("%T\n", v)
	case person:
    fmt.Println(v)
	default:
	fmt.Printf("%T\n", v)
	}
}
type myexpenses interface {
	mycost() float64
}
type myemails struct {
	isSubscribed bool
	body string
	toAddress string
}
type mysmss struct {
	isSubscribed bool
	body string
	toPhoneNumber string
}
type myinvalid struct{}
func getExpenseReportTypeSwitch(e myexpenses) (string, float64) {
switch v := e.(type) {
case myemails:
	return v.toAddress, v.mycost()
case mysmss:
	return v.toPhoneNumber, v.mycost()
default:
    return "", 0.0
}
}
func (e myemails) mycost() float64 {
if !e.isSubscribed {
	mycost := 0.05 * float64(len(e.body))
	return  mycost
}
	mycost := 0.01 * float64(len(e.body))
	return  mycost
	
}
func (s mysmss) mycost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}
func (i invalid) mycost() float64 {
	return 0.0
}
func mymyTest(e myexpenses){
	address, cost := getExpenseReportTypeSwitch(e)
	switch e.(type) {
	case myemails:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("=================================================")
	case mysmss:
		fmt.Printf("Report: The smss going to %s will cost: %.2f\n", address, cost)
		fmt.Println("=================================================")
	default:
        fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}
func TypeSwitches() {
printNumbericValue(1)
// prints int
printNumbericValue("1")
// prints string
printNumbericValue(struct{}{}) // this gets initialized with no types and no default values
// prints "struct {}"
printNumbericValue(person{ name: "Gideon Babalola"})
mymyTest(myemails{
	isSubscribed: true,
	body: "hello there",
	toAddress: "john@doe.com",
})
mymyTest(myemails{
	isSubscribed: false,
	body: "This meeting could have been an email",
	toAddress: "jane@doe.com",
})
mymyTest(myemails{
	isSubscribed: false,
	body: "This meeting could have been an email",
	toAddress: "elan@doe.com",
})
mymyTest(mysmss{
	isSubscribed: false,
	body: "This meeting could have been an email",
	toPhoneNumber: "+155555509832",
})
mymyTest(mysmss{
	isSubscribed: false,
	body: "This meeting could have been an email",
	toPhoneNumber: "+155555504444",
})
mymyTest(invalid{})
}

//Rule of thumb, interfaces should not be aware of their underlying type
// Interfaces are not classes
// 1) Interfaces are not classes, they are slimmer
// 2) Interfaces don't have constructors or deconstructors that require that data is created or destroyed
// 3) Interfaces are not hierarchical by nature, though there is syntactic sugar to create interfaces that happens to be supersets of other interfaces
// 4) Interfaces defines function signatures, but not underlying behaviour
// Making an interface often won't DRY up your code in regards to struct methods. For example, if five types satisfy the fmt.Stringer interface, they all need their own version of the String function.