package main

import "fmt"
import "strings"

func wordListing(num int, word string) string {
	var word_list string
	switch num {

	case 1, 5, 6, 7, 8, 9, 15, 16, 19:
		word_list = word[:1]
	default:
		word_list = word[:2]
	}
	return word_list
}

func main() {
	sentence := []rune("Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.")

	word_list := make(map[string]int)
	var word []string
	letter_count := 0

	for _, v := range sentence {
		if string(v) != " " && string(v) != "," && string(v) != "." {
			word = append(word, string(v))
		} else {
			if len(word) == 0 {
				continue
			}

			letter_count++
			bundled_word := strings.Join(word, "")
			word = nil

			//bundled_wordを加工してマップに格納
			word_list[wordListing(letter_count, bundled_word)] = letter_count
		}
	}

	fmt.Println(word_list)
}
