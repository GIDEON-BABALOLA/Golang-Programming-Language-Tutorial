package main
import (
	"errors"
	"fmt"
)
// The Go Standard library provides an "errors" package that makes it easy to deal with errors.
// func ErrorsPackage(){
// var err error = errors.New("something went wrong")
// } 
type divideErrorP struct {
	dividend float64
}
func (de divideErrorP) Error() string {
return fmt.Sprintf("can not divide %v by zero",
 de.dividend,
)
}
func divideP(dividend, divisor float64) (float64, error){
fmt.Printf("Dividing %.2f by %.2f ... \n", dividend, divisor)
if divisor == 0 {
	return dividend / divisor, errors.New("no dividing by 0")
}
   return dividend / divisor, nil
}
func testDivideP(dividend, divisor float64){
	result, err := divideP(dividend, divisor)
	if err != nil {
		divError := err.Error()
        fmt.Println(divError)
		fmt.Println("==============================") 
		return
	}
	fmt.Printf("Quotient: %.2f \n", result)
	fmt.Println("==================================")
}
func ErrorsPackage(){
testDivideP(10, 0)
testDivideP(10, 2)
testDivideP(15, 30)
testDivideP(6, 3)
} 