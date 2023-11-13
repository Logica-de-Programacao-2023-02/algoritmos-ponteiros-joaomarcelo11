package main

import "fmt"

func updateIntValue(ptr *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*ptr = sum
}

func main() {
	var num int
	updateIntValue(&num, 5)
	fmt.Println("Updated value:", num) // Saída: Updated value: 15
}