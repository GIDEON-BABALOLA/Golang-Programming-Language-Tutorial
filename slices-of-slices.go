package main

import "fmt"
func createMatrixSlice(rows, cols int) [][]int {
 matrix := [][]int{}
 for i := 0; i < rows; i++ {
	row := []int{}
	for j := 0; j < cols; j++ {
		row = append(row, i*j)
	}
	matrix = append(matrix, row)
 }
 return matrix
}
func testMatrixSlice(rows, cols int){
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrixSlice(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("====== END REPORT =======")
}
func SlicesOfSlices(){
testMatrixSlice(5, 5)
}