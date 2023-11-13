package main

import "fmt"

func reverseString(ptr *string) {
	str := *ptr
	runes := []rune(str)
	n := len(runes)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*ptr = string(runes)
}

func main() {
	text := "Hello, World!"
	reverseString(&text)
	fmt.Println("Reversed string:", text) // Saída: Reversed string: !dlroW ,olleH
}