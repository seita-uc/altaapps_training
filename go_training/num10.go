package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func wordsSplit(s string) []string {
	words := strings.Split(s, " ")
	return words
}

func randomizeWord(word string) string {
	var newWord string
	runes := []rune(word)
	length := len(runes)

	if length > 4 {
		newRunes := make([]rune, length)
		randIdx := rand.Perm(length - 2)
		newRunes[0] = runes[0]
		newRunes[length-1] = runes[length-1]

		for i := range randIdx {
			//randIdxの各要素に１を足せば最初と最後以外のインデックスがとれる
			newRunes[i+1] = runes[randIdx[i]+1]
		}
		newWord = string(newRunes)
	} else {
		newWord = word
	}
	return newWord
}

func randomizeSent(sent string) string {
	words := wordsSplit(sent)

	for i := 0; i < len(words); i++ {
		words[i] = randomizeWord(words[i])
	}
	return strings.Join(words, " ")
}

func main() {

	sent := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	randomSent := randomizeSent(sent)

	fmt.Printf("sent : %v\n", sent)       //元の文
	fmt.Printf("sent : %v\n", randomSent) //ランダム化された文
}
