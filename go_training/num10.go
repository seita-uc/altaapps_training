package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func WordsSplit(s string) []string {
	words := strings.Split(s, " ")
	return words
}

func RandomizeWord(word string) string {
	var new_word string
	runes := []rune(word)
	length := len(runes)

	if length > 4 {
		new_runes := make([]rune, length)
		rand_idx := rand.Perm(length - 2)
		new_runes[0] = runes[0]
		new_runes[length-1] = runes[length-1]

		for i := range rand_idx {
			//rand_idxの各要素に１を足せば最初と最後以外のインデックスがとれる
			new_runes[i+1] = runes[rand_idx[i]+1]
		}
		new_word = string(new_runes)
	} else {
		new_word = word
	}
	return new_word
}

func RandomizeSent(sent string) string {
	words := WordsSplit(sent)

	for i := 0; i < len(words); i++ {
		words[i] = RandomizeWord(words[i])
	}
	return strings.Join(words, " ")
}

func main() {

	sent := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	random_sent := RandomizeSent(sent)

	fmt.Printf("sent : %v\n", sent)        //元の文
	fmt.Printf("sent : %v\n", random_sent) //ランダム化された文
}
