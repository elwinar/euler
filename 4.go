// A palindromic number reads the same both ways. The largest palindrome
// made from the product of two 2-digit numbers is 9009 = 91 × 99.
// 
// Find the largest palindrome made from the product of two 3-digit 
// numbers.
package main

func main() {
	for p := uint64(1000*1000); p > uint64(100*100); p-- {
		if isPalindrome(p) {
			for i := uint64(100); i < uint64(1000); i++ {
				if p % i == 0 && p / i < 1000 {
					println(p)
					return
				}
			}
		}
	}
}

func isPalindrome(number uint64) bool {
	var n = number
	var mirror = uint64(0)
	for n > 0 {
		mirror = mirror * 10 + n % 10
		n = n / 10
	}
	return mirror == number
}
