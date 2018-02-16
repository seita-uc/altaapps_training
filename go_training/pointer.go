package main

import (
	"fmt"
)

func main() {
	var n int = 100

	pointer := &n

	fmt.Println("address of n:", &n)
	fmt.Println("value of pointer:", pointer)

	fmt.Println("value of n:", n)

	fmt.Println("content of pointer:", *pointer)

}
