package main

import "fmt"
import "strings"

func main() {
	str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."

	word_list := make(map[string]int)
	var word []string
	var bundled_word []string
	num := 0
	for i := 1; i <= len(str); i++ {
		if str[i-1:i] != " " && str[i-1:i] != "," && str[i-1:i] != "." {
			word = append(word, str[i-1:i])
		} else {
			if len(word) == 0 {
				continue
			}
			num++
			bundled_word = strings.Join(word, "")
			fmt.Printf("%v   %v   %v\n", bundled_word, len(bundled_word), num)
			word = nil
			//			switch num {
			//	case 1, 5, 6, 7, 8, 9, 15, 16, 19:
			//	word_list[bundled_word[:2]] = num
			//	default:
			//		word_list[bundled_word[:3]] = num

			//}
			if num == 1 {
				word_list[bundled_word[:2]] = num
				fmt.Println(word_list)
				//	} else {

				//		word_list[bundled_word[:3]] = num
				//		fmt.Println(word_list)
			}

		}
	}

	//	fmt.Println(word_list)

}
