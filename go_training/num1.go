package main

import "fmt"

func main() {
	//問題01
	str := "stressed"
	for i := len(str) - 1; i >= 0; i -= 1 {
		fmt.Print(str[i])
	}
}
