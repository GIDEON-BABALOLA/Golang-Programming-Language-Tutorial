package main
import "fmt"
type divideError struct {
	dividend float64
}
func (de divideError)  Error() string {
return fmt.Sprintf("can not divide %v by zero",
 de.dividend,
)
}
func divide(dividend, divisor float64) (float64, error){
fmt.Printf("Dividing %.2f by %.2f ... \n", dividend, divisor)
if divisor == 0 {
	return dividend / divisor, divideError{dividend: dividend}
}
   return dividend / divisor, nil
}
func testDivide(dividend, divisor float64){
	result, err := divide(dividend, divisor)
	if err != nil {
		divError := err.Error()
        fmt.Println(divError)
		fmt.Println("==============================")
		return
	}
	fmt.Printf("Quotient: %.2f \n", result)
	fmt.Println("==================================")
}
func CustomError(){
testDivide(10, 0)
testDivide(10, 2)
testDivide(15, 30)
testDivide(6, 3)
}