// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// 
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
package main

func main() {
	var primes = [...]uint64{2, 3, 5, 7, 11, 13, 17, 19}
	var number uint64 = 1
	
	for _, p := range primes {
		var factor uint64 = 1
		for factor <= 20 {
			factor *= p
		}
		number *= factor / p
	}
	
	println(number)
}
