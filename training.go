package main

import "fmt"

func main() {
	str := "stressed"
	for i := len(str); i >= 1; i -= 1 {
		fmt.Printf(str[i-1 : i])
	}

}
