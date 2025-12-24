package main
// A type can represent multiple interfaces, it just needs to have the mthods for all of the different interfaces
// A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.
import "fmt"
func (e email) cost() float64 {
if !e.isSubscribed {
	cost := 0.05 * float64(len(e.body))
	return  cost
}
	cost := 0.01 * float64(len(e.body))
	return  cost
	
}
func (e email) print()  {
fmt.Println(e.body)
}
type expense interface {
	cost() float64
}
type printer interface {
	print()
}
type email struct {
	isSubscribed bool
	body string
}
func print(p printer) {
	p.print()
}
func testHere(e expense, p printer){
fmt.Printf("Printing with cost: $%.2f .....\n", e.cost())
p.print()
fmt.Println("==========================================")
}

// Name your interface arguments
type Copier interface {
	Copy(string, string) int
}
type NewCopier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
func MultipleInterface(){
	print(email{
		isSubscribed: true,
		body: "This user is subscribed o",
	})
	testHere(email{
	    isSubscribed: true,
		body: "This user is subscribed o",
	}, email{
		isSubscribed: false,
		body: "This user is not subscribed o",
	})
	e := email{
		isSubscribed: true,
		body: "hello there",
	}
	testHere(e, e)
	e = email{
		isSubscribed: false,
		body: "I want my money back",
	}
	testHere(e, e)
}