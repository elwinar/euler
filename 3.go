// The prime factors of 13195 are 5, 7, 13 and 29.
// 
// What is the largest prime factor of the number 600851475143 ?
package main

func main() {
	var number uint64 = 600851475143
	
	var factor uint64 = 2
	for factor != number {
		for number % factor == 0 {
			number = number / factor
		}
		factor++
	}
	
	println(factor)
}
