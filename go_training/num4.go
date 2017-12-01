package main

import "fmt"
import "strings"

func main() {
	sentence := []rune("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")
	var word []string
	for _, letter := range sentence {
		if string(letter) != " " && string(letter) != "," && string(letter) != "." {
			word = append(word, string(letter))
		} else {
			if len(word) == 0 {
				continue
			}
			fmt.Printf("%v   %v\n", strings.Join(word, ""), len(word))
			word = nil
		}
	}
}
