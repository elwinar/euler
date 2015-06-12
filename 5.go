// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// 
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
package main

func main() {
	var primes = [...]uint64{2, 3, 5, 7, 11, 13, 17, 19}
	var factors = make(map[uint64]uint64)
	
	for n := uint64(2); n <= 20; n++ {
		number := n
		var composition = make(map[uint64]uint64)
		for _, p := range primes {
			for number % p == 0 {
				composition[p]++
				number /= p
			}
		}
		
		for k, v := range composition {
			if v > factors[k] {
				factors[k] = v
			}
		}
	}
	
	var sum uint64 = 1
	for k, v := range factors {
		for i := uint64(0); i < v; i++ {
			sum *= k
		}
	}
	
	println(sum)
}
