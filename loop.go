package main

import "fmt"
// The basic loop in Go is written in standard C-like syntax
func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		// totalCost += 1.0 + (0.01 * float64(i))
		totalCost = totalCost + (1.0 + (0.01 * float64(i)))
	} 
	return totalCost
}
// Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loops to run forever.
func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost = totalCost + (1.0 + (0.01 * float64(i)))
		if totalCost > thresh {
			return i
		}
	}
}
// THERE IS NO WHILE LOOP IN GO
// Most programming languages have a concept of a while loop. Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a condition
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	for actualCostInPennies <= float64(maxCostInPennies){
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMessagesToSend
}
func testLoops(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===================================================")
}
func testOmitLoops(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("==================================================")
}
func Loops(){
testLoops(10)
testLoops(20)
testLoops(30)
testLoops(40)
testOmitLoops(10)
testOmitLoops(20)
testOmitLoops(30)
testOmitLoops(40)
}
// 