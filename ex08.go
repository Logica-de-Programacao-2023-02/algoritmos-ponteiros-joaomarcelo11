package main

import (
	"fmt"
)

func modifyValue(ptr *int) {
	*ptr = 42
}

func main() {
	ptr := new(int)
	*ptr = 10

	fmt.Println("Valor antes da modificação:", *ptr)

	modifyValue(ptr)

	fmt.Println("Valor após a modificação:", *ptr) 
	free(ptr)
}

func free(ptr *int) {
	ptr = nil
}