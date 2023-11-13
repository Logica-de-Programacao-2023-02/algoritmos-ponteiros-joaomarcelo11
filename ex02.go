package main

import "fmt"

func updateIntValue(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0 
	} else {
		*ptr = 1 
	}
}

func main() {
	num := 7
	updateIntValue(&num)
	fmt.Println("Updated value:", num) 
}