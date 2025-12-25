package main
import "fmt"
// A type switch makes it easy to do several assertions in a series
// A type switch is similar to a regular switch statement, but the cases specify types instead of values
// The reason why we can cast any of the types int, and string to interface num to is because an empty interface in Go is implemented by every types
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
func TypeSwitches() {
printNumbericValue(1)
// prints int
printNumbericValue("1")
// prints string
printNumbericValue(struct{}{}) // this gets initialized with no types and no default values
// prints "struct {}"
printNumbericValue(person{ name: "Gideon Babalola"})
}