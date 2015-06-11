// If we list all the natural numbers below 10 that are multiples of 3
// or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
package main

// Computationnal solution: the sum of all numbers from 1 to N
// can be computed with certainty with the N(N+1)/2 formula.
// This formula is still usable with a step S, by computing for a N/S
// range then multiplying by S.
// For all numbers multiples of 3 and 5, we compute the sum of all 
// 3-multiples and 5-multiples, then subtract the (3*5)-multiples, which
// would otherwise be counted twice.
func main() {
	println(sum(999, 3) + sum(999, 5) - sum(999, 15))
}

func sum(n int, s int) int {
	p := n / s
	return s * p * (p + 1) / 2
}
