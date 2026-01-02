package main

import (
	"fmt"
	"math"
)

// Go supports the standard modulo operator
// 7 % 3 returns 1
// Logical AND operator
// true && false // false
// true && true // true
// Logical OR operator
// true || false // true
// false || false // false


// Complete the fizzbuzz function that prints the numbers 1 to 100 inclusive each of their own line, but substitutes multiples of 3 for the text fizz and multiples of 5 for buzz. For multiples of 3 AND 5 print instead fizzbuzz


// Sudocode For Prime numbers
// for n in range(2, max+1):
//   if n is 2:
//.    n is prime, print it
//.  if n is even:
//.    n is not prime, skip to the next n
//. for i in range (3, sqrt(n) + 1):
//     if i can be multiplied into n:
//.       n is not prime, skip to the next n
//.  n is prime, print it

// Breakdown of printPrimes sudocode
// We skip even numbers because they can't be prime
// We only check up to the square root because anything higher than the square root has no chance of multiplying evenly into n
// We start checking at 2 because 1 is not prime
func printPrimes(max int) {
	for n := 2; n < max + 1; n++{
		if n == 2 {
			fmt.Println(n)
		}
		if n % 2 == 0 {
			continue
		}
		isPrime := true
	    for i := 3.0; i < math.Sqrt(float64(n)) + 1 ; i++{
        if n % int(i) == 0 {
			isPrime = false
			break
		}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}
func testPrime(max int) {
	fmt.Printf("Primes up to %v\n", max)
	printPrimes(max)
	fmt.Println("================================")
}
func Fizzbuzz() {
	for i := 1; i <=100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
// Continue
// The continue keyword stops the current iteration of a loop and continues to the next iteration. continue is a powerful way to use the "guard clause" pattern within loops.
for i :=0; i < 10; i++ {
	if i % 2 == 0 {
		continue
	}
	fmt.Println(i)
    // 1
    // 3
    // 5
    // 7
    // 9
}
// Break
// The break keyword stops the current iteration of a loop and exits the loop
for i := 0; i < 10; i++ {
	if i == 5 {
		break
	}
	fmt.Println(i)
}
// 0
// 1
// 2
// 3
// 4 
testPrime(20)
}