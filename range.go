package main

import "fmt"

// RANGE
// Go provides syntactic sugar to iterate easily over elements of a slice:
func indexOfFirstBadWord(msg []string, badWords []string) int {
   for m, message := range msg {
     for _, badWord := range badWords {
      if badWord == message {
	  return m
	  }
	 }
   }
   return -1
}
func testRange(msg []string, badWords []string){
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("=============================================")
}
func Range(){
// fruits := []string{"apple", "banana", "grape"}
// for i, fruit := range fruits {
// 	fmt.Println(i, fruit)
// }
badWords := []string{"crap", "shoot", "dang", "frick"}
message :=  []string{"hey", "there", "john"}
testRange(message, badWords)

message =  []string{"ugh", "oh", "my", "frick"}
testRange(message, badWords)

message = []string{"what", "the", "shoot", "I", "hate", "that"}
testRange(message, badWords)
}