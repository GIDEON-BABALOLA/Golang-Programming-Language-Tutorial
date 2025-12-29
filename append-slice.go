package main

// APPEND
// The built-in append function is used to dynamically add elements to a slice:
// The append function is a variadic function under the hood
// func appen(slice []Type, elems ...Type) []Type
// If the underlying array is not large enough, append() will create a new underlying array and point the slice to it.
// Notice that append() is variadic, the following are all valid:
// slice = append(slice, oneThing)
// slice = append(slice, firstThing, secondThing)
// slice = append(slice, anotherSlice...)
type cost struct {
	day int
	value float64
}
func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
costsByDay = append(costsByDay, 0.0)
		}
// costsByDay = append(costsByDay, cost.value)
costsByDay[cost.day] = costsByDay[cost.day] + cost.value
	}
	return costsByDay
}
func AppendSlice() {

}
