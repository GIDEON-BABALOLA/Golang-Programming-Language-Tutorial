package main
import "fmt"
// NESTED MAPS
// Maps can contain maps, creating a nested structure. For example:
// map[string]map[string]int
// map[rune]map[string]int
// map[int]map[string]map[string]int
func getNamesCounts( names []string) map[rune]map[string]int {
myNestedMap := make(map[rune]map[string]int)
for _, name := range names {
	if name == "" {
		continue
	}
	firstKey := rune(name[0])
	secondKey := name
	count := myNestedMap[firstKey][secondKey]
	count++
	myNestedMap[firstKey][secondKey] = count
}
return myNestedMap
}
func testNestedMaps(names []string, initial rune, name string){
	fmt.Printf("Generating counts for %v names...\n", len(names))
	nameCounts := getNamesCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
	fmt.Println("==============================")
}
func NestedMaps(){
testNestedMaps(getNamesNestedMap(50),  'M', "Matthew")
testNestedMaps(getNamesNestedMap(100), 'G', "George")
testNestedMaps(getNamesNestedMap(150), 'D', "Drew")
testNestedMaps(getNamesNestedMap(200), 'P', "Philip")
testNestedMaps(getNamesNestedMap(250), 'B', "Bryant")
testNestedMaps(getNamesNestedMap(300), 'M', "Matthew")
}
func getNamesNestedMap(length int) []string {
	names := []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew",
	}
	if length > len(names){
		length = len(names)
	}
	return names
}
// A function that accepts a function as a paramter or a function that returns a function as its return value is known as an higher-order function