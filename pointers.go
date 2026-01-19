package main
import "fmt"
// As we have learned, a variable is a named location in memory that stores a value. We can manipulate the value of the variable by assigning a new value to it or by performing operations on it. When we assign a value to a variable, we are storing that value in a specific location in memory.
// A Pointer Is A Variable
// A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored not the data itself
// The * syntax defines a pointer:
// var p *int
// The & operator generates a pointer to its operand.
// myString := "hello"
// myStringpTR = &myString
// The zero value for a pointer is nil
// Two reasons we should use a pointer
// 1) You want to be able to pass a value into a function and change the value and have those changes persist outside of the function.
// 2) If you are very concerned about the performance of your program, because everytime you have to create the copy of some variable in memory, then copying takes some time. You can make micro optimizations with pointers if you want to avoid all of that memory copying.
func Pointers() {

}