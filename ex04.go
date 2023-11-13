package main

import "fmt"

func updateIntValue(ptr *int) {
	num := *ptr
	lastDigit := num % 10
	num /= 10
	secondLastDigit := num % 10
	sum := lastDigit + secondLastDigit
	*ptr = sum
}

func main() {
	num := 1234
	updateIntValue(&num)
	fmt.Println("Updated value:", num) // Saída: Updated value: 7
}