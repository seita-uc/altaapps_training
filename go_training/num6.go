package main

import (
	"fmt"
	"strings"
)

func sentenceIntoWords(sentence string) []string {
	split_str := []rune(sentence)
	var word []string
	var splited_words []string
	for i, v := range split_str {
		if string(v) != " " && string(v) != "," && string(v) != "." {
			word = append(word, string(v))
			if i == len(split_str)-1 {
				splited_words = append(splited_words, strings.Join(word, ""))
			}
		} else {
			if len(word) == 0 {
				continue
			}
			splited_words = append(splited_words, strings.Join(word, ""))
			word = nil
		}
	}
	return splited_words
}

func sentenceIntoLetters(sentence string) []string {
	split_str := []rune(sentence)
	var letters []string
	for _, v := range split_str {
		if string(v) != " " && string(v) != "," && string(v) != "." {
			letters = append(letters, string(v))
		}
	}
	return letters
}

func createNgram(w_ele []string, l_ele []string) ([]string, []string) {

	//単語bi-gram
	var word_result []string
	for i := 0; i < len(w_ele); i++ {
		if i+1 == len(w_ele) {
			break
		}
		word_result = append(word_result, strings.Join(w_ele[i:i+2], "-"))
	}

	//文字bi-gram
	var letter_result []string
	for i := 0; i < len(l_ele); i++ {
		if i+1 == len(l_ele) {
			break
		}
		letter_result = append(letter_result, strings.Join(l_ele[i:i+2], "-"))
	}

	return word_result, letter_result
}

func main() {

	fmt.Println(createNgram(sentenceIntoWords("I am an NLPer"), sentenceIntoLetters("I am an NLPer")))
}
