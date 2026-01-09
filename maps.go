package main

import (
	"errors"
	"fmt"
)

// Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.
// The zero value of a map is nil just like the zero value of an array is nil.
// We can create a map by using a literal or by using the make() function

func getUserMap(names []string, phoneNumbers []int) (map[string]userMap, error) {
if len(names) != len(phoneNumbers) {
	return nil, errors.New("Invalid Sizes")
}
newMap := make(map[string]userMap)
for i := 0; i < len(names); i++{
key := names[i]
value := phoneNumbers[i]
newMap[key] = userMap{ name: key, phoneNumber: value}
}
return newMap, nil
}
type userMap struct{
	name string
	phoneNumber int
}
func testMaps(names []string, phoneNumbers []int){
	fmt.Println("Creating map...")
	defer fmt.Println("=============================")
	users, err := getUserMap(names, phoneNumbers)
	fmt.Println(users)
	if err != nil {
		fmt.Println(err)
		return;
	}
	for _, name := range names {
		fmt.Printf("Key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}
func Maps() {
ages := make(map[string]int)
ages["john"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24
ages = map[string]int{
	"john": 37,
	"Mary": 21,
}
// The len() function works on a map, it returns the total number of key/value pairs
// fmt.Println(len(ages))
// names := []string{"Gideon", "Babalola"}
// phoneNumbers := []int{9, 10, 11, 12}
// testMaps(names, phoneNumbers)
testMaps(
	[]string{"John", "Bob", "Jill"},
	[]int{144444444, 23344444444, 33333333},
)
testMaps(
	[]string{"John", "Bob"},
	[]int{1444449944444, 233444488494444, 333490333333},
)
}


// func getUserMap(names []string, phoneNumbers []int) (map[string]int, error) {
// if len(names) != len(phoneNumbers) {
// 	return nil, errors.New("Invalid Sizes")
// }
// newMap := make(map[string]int)
// for _, name := range names {
// 	for _, number := range phoneNumbers {
// 		newMap[name] = number
// 		break;
// 	}
// }
// return newMap, nil
// }
// So basically what I wanted to accomplish was to directly map items from one slice to another slice using a map data structure in Go. So this was it, I had two slices, the first one an slice of names
//  {"John", "Bob", "Jill"}
// And the second one an slice of phone numbers
// {081149787227, 091076478721", "0818060467636"}
// And my desired result was to have a map like this
// {john: +2348149787227, Bob : +2349076478721, Jill : +2348060467636}
// So initially my normal self would just directly start writing the code and my initial approach was to use a nested loop.