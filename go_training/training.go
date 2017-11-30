package main

import "fmt"

func main() {
	{ //問題01
		str := "stressed"
		for i := len(str); i >= 1; i -= 1 {
			fmt.Printf(str[i-1 : i])
		}
		fmt.Printf("\n")
	}
	{ //問題02
		str := []rune("パタトクカシーー")
		fmt.Printf("%c%c%c%c\n", str[1], str[3], str[5], str[7])
		//extractStr := str[1] + str[3] + str[5] + str[7]
		//for i := range extractStr {
		//	fmt.Printf("%c", extractStr[i])
		//}
		//文字の連結ができない
	}
	{ //問題03
		str1 := []rune("パトカー")
		str2 := []rune("タクシー")
		for num := 0; num < 4; num++ {
			fmt.Printf("%c%c", str1[num], str2[num])

		}

	}
}
