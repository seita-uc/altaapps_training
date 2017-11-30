package main

import "fmt"
import "strings"

func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	var word []string
	for i := 1; i <= len(str); i++ {
		if str[i-1:i] != " " && str[i-1:i] != "," && str[i-1:i] != "." {
			word = append(word, str[i-1:i])
		} else {
			if len(word) == 0 {
				continue
			}

			fmt.Printf("%v   %v\n", strings.Join(word, ""), len(word))
			word = nil
		}
	}
}
