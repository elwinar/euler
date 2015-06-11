// The prime factors of 13195 are 5, 7, 13 and 29.
// 
// What is the largest prime factor of the number 600851475143 ?
package main

import "math"

// Find all the integer factors of the Big Evil Number, then feed them
// to a concurrent prime sieve to extract the primes ones.
func main() {
	var number uint64 = 600851475143
	var root uint64 = uint64(math.Sqrt(float64(number)))
	
	var factors []uint64
	for i := uint64(2); i < root; i++ {
		if number % i == 0 {
			factors = append(factors, i)
		}
	}
	
	var out = make(chan uint64)
	go func(out chan uint64) {
		for _, f := range factors {
			out <- f
		}
		close(out)
	}(out)
	
	var factor uint64 = 0
	for {
		p, ok := <- out
		
		if !ok {
			break
		}
		
		factor = p
		
		in := out
		out = make(chan uint64)
		
		go func(f uint64, in chan uint64, out chan uint64) {
			for n := range in {
				if n % f != 0 {
					out <- n
				}
			}
			close(out)
		}(p, in, out)
	}
	
	println(factor)
}
