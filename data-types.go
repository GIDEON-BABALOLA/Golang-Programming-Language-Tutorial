package main
import "fmt"
func dataTypes() {
	congrats := "Happy Birthday"
	penniesPerText := 2.0
	// Declaring text on the same line
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
	fmt.Println(congrats)
	fmt.Println(averageOpenRate, displayMessage)
	accountAge := 2.6
	accountAgeInt := int(accountAge)
	// Constants in Go does not support short declaration syntax such as :=
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"
	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
	// Constants in Go can only be computed at compile time but not at runtime, so if the inputs to a constant are unknown, an error will be thrown at compile time
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = minutesInHour * secondsInMinute
	fmt.Println("number of seconds in an hour:", secondsInHour)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}