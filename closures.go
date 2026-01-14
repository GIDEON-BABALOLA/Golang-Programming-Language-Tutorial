package main
import "fmt"
// CLOSURES
// A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.
// In this example, the concatter() function returns a function that has a reference to an enclosed doc value. Each successive call to harryPotterAggregator mutates the same doc variable.
func concatter() func(string) string {
	doc := ""
	return func(word string) string{
		doc = doc + word + " "
		return doc
	}
}
func adder() func(int) int{
	sum := 0;
	return func(input int) int {
		sum = sum + input
		return sum
	}
}
type emailBill struct {
	costInPennies int
}
func testClosures(bills []emailBill){
	defer fmt.Println("==================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}
func Closures(){
harryPotterAggregator := concatter()
harryPotterAggregator("Mr.")
harryPotterAggregator("and")  
harryPotterAggregator("Mrs.")
harryPotterAggregator("Dursley")
harryPotterAggregator("of")
harryPotterAggregator("number")
harryPotterAggregator("four,")
harryPotterAggregator("Privet")
fmt.Println(harryPotterAggregator("Drive"))
// Mr. and Mrs. Dursley of number four, Privet Drive 
testClosures([]emailBill{
	{costInPennies: 45},
	{costInPennies: 32},
	{costInPennies: 43},
	{costInPennies: 12},
	{costInPennies: 34},
	{costInPennies: 54},
})
// so if an struct only has one attribute, we dont have to include the attribute when assigning to the struct
testClosures([]emailBill{
	{12},
	{12},
	{976},
	{12},
    {543},
})
testClosures([]emailBill{
	{12},
	{12},
	{976},
	{12},
    {543},
})
testClosures([]emailBill{
	{743},
	{13},
	{8},
})
}
// When a variable is enclosed in a closure, the enclosing function has access to a mutable reference of the original value