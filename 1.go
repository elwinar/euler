// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6
// and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
package main

//import "math"

func main() {
	println(sum(1000, 3) + sum(1000, 5) - sum(1000, 15))
}

func sum(n int, s int) int {
	return s * ((n - 1) / s) * (n / s) / 2
}

