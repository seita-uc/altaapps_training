package main

import "fmt"

func main() {
	str := []rune("stressed")
	for i := len(str) - 1; i >= 0; i -= 1 {
		fmt.Printf("%c", str[i])
	}
}
