// The sum of the squares of the first ten natural numbers is,
// 12 + 22 + ... + 102 = 385
//
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)2 = 552 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
package main

func Brute(n int) int {
	sum := 0
	sumsquares := 0
	for i := 0; i <= n; i++ {
		sum += i
		sumsquares += i * i
	}
	return sum*sum - sumsquares
}

func Elegant(n int) int {
	sum := n * (n + 1) / 2
	sumsquare := (2*n + 1) * (n + 1) * n / 6
	return sum*sum - sumsquare
}

func main() {
	println(Brute(100))
	println(Elegant(100))
}
