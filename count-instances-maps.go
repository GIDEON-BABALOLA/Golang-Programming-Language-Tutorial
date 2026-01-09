package main

import (
	// "crypto/md5"
	"crypto/md5"
	"fmt"
	"io"
	// "io"
)
func getCounts( userIDs []string) map[string]int {
	// ?
	counts := make(map[string]int)
	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}
func testCountInstances( userIDs []string, ids []string){
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))
	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)
	}
	fmt.Println("=========================================")
}
func CountInstances(){
	userIds := []string{}
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIds = append(userIds, key[:2])
	}
	testCountInstances(userIds, []string{"00", "ff", "dd"})
	testCountInstances(userIds, []string{"aa", "12", "32"})
	testCountInstances(userIds, []string{"bb", "33",})
}

// userIDs ["2eue", "uri8", "iru8", "ijr9", "uri9", "uit8"]
// Attempting to get a value from a map that does not exist will make your code panic