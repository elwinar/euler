// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?
package main

import "fmt"

func Brute(n int64) int64 {
	var primes []int64
	var i int64 = 2
	for len(primes) < 10001 {
		prime := true
		for _, p := range primes {
			if i%p == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, i)
		}
		i++
	}
	return primes[n-1]
}

func main() {
	fmt.Println(Brute(10001))
}
