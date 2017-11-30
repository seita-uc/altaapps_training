package main

import "fmt"

func main() {
	//問題03
	str1 := []rune("パトカー")
	str2 := []rune("タクシー")
	for num := 0; num < 4; num++ {
		fmt.Printf("%c%c", str1[num], str2[num])

	}
}
