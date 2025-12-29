package main
// import "fmt"
// MAKE
// Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:
func MakeSlices(){
// func make([]T, len, cap) []T
// mySlice := make([]int, 5, 10)
// the capacity argument is usually omitted and defaults to the length of the slice
// mySlice = make([]int, 5)
// slices created with make will be filled with the zero value of the type
// If we want to create a slice with a specific set of values, we can use a slice literal
// myNewSlice := []string{"I", "love", "go"}
// Note that the array brackets do not have 3 in them. If they did, you'd have an array instead of a slice
// The cap() function means the maximum length of the slice before reallocation of the array is necessary
// The len() function returns the current length of the slice
// The len() and cap() returns 0 when your slice is nil 
}