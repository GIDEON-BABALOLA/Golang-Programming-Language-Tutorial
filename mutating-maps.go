package main

import (
	"errors"
	"fmt"
)

// m[key] = elem // Insert an element
// elem = m[key] // Get an element
// delete(m, key) // Delete an element
// elem, ok := m[key] // Checks if an element exists, So if key is in m, then ok is true. If not, ok is false. If the key is not in the map, then elem is the zero value for the map's element type.
// Note On Passing Maps
// Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy

// What maks a type qualify to be able to be used as a map key, Any type can be used as a map value but not every type can be used as a map key and that is because map key may be of any type that is comparable, so we cannot use slices, maps and functions as our map keys
type mutateUsers struct {
	name string
    number int
	scheduledForDeletion bool
}
func deleteIfNecessary(users map[string]mutateUsers, name string) (deleted bool, err error){
	exists, ok := users[name]
	if !ok {
return false, errors.New("not found")
	}
    if !exists.scheduledForDeletion {
return false, nil
	}   
delete(users, name)
return true, nil
}
func testMutating(users map[string]mutateUsers, name string){
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("=====================")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted:", name)
		return
	}
	fmt.Println("Did not delete:", name)

}
func MutatingMaps() {
users := map[string]mutateUsers{
	"john": {
		name: "john",
		number: 34848848484,
		scheduledForDeletion: true,
	},
	"elon" : {
	    name: "elon",
		number: 2838393939,
		scheduledForDeletion: true,
	},
	"breanna": {
	    name: "breanna",
		number: 393939393993934,
		scheduledForDeletion: false,
	},
}
testMutating(users, "john")
testMutating(users, "musk")
}