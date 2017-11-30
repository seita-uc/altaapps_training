package main

import "fmt"

func main() {
	//問題02
	str := []rune("パタトクカシーー")
	fmt.Printf("%c%c%c%c\n", str[1], str[3], str[5], str[7])
	//extractStr := str[1] + str[3] + str[5] + str[7]
	//for i := range extractStr {
	//  fmt.Printf("%c", extractStr[i])
	//}
	//文字の連結ができない

}
