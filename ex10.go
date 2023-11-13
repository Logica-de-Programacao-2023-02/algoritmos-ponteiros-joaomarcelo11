package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func fillPrimes(slice *[]int, size int) {
	count := 0
	num := 2

	for count < size {
		if isPrime(num) {
			*slice = append(*slice, num)
			count++
		}
		num++
	}
}

func main() {
	var primes []int
	size := 10

	fillPrimes(&primes, size)

	fmt.Println("Primeiros", size, "números primos:", primes)
}