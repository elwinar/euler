// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
package main

func main() {
	println(sum(999, 3) + sum(999, 5) - sum(999, 15))
}

func sum(n int, s int) int {
	p := n / s
	return s * p * (p + 1) / 2
}
