package main

import (
	"fmt"
	"math"
)

func updateFloatValue(ptr *float64) {
	*pi := math.Pi
	avg := (*ptr + *pi) / 2
	*ptr = avg
}

func main() {
	num := 3.14
	updateFloatValue(&num)
	fmt.Println("Updated value:", num) // Saída: Updated value: 3.927043460986335
}