package main

import "fmt"
import "strings"

func createNgram(sentence string) []string {

	letters := strings.Split(sentence, "")

	//文字bi-gram
	var letterResult []string
	for i := 0; i < len(letters); i++ {
		if i+1 == len(letters) {
			break
		}
		letterResult = append(letterResult, strings.Join(letters[i:i+2], ""))
	}
	return removeDuplication(letterResult)
}

func remove(search string, set []string) []string {
	result := []string{}
	for _, v := range set {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

func include(s string, set []string) bool {
	for _, v := range set {
		if v == s {
			return true
		}
	}
	return false
}

func union(set1 []string, set2 []string) []string {
	new_str := []string{}
	for _, v := range set1 {
		if !include(v, set2) {
			new_str = append(new_str, v)
		}
	}
	new_str = append(new_str, set2...)
	return new_str
}

func removeDuplication(set []string) []string {
	m := make(map[string]bool)
	new_str := []string{}
	for _, v := range set {
		if !m[v] {
			m[v] = true
			new_str = append(new_str, v)
		}
	}
	return new_str
}

func intersection(set1 []string, set2 []string) []string {
	num := 0
	var intersect_str []string
	for _, v := range set1 {
		for i, s := range set2 {
			if v == s {
				num++
			}
			if i == len(set2) {
				break
			}
		}
		if num == 1 {
			intersect_str = append(intersect_str, v)
		}
		num = 0
	}
	return intersect_str
}

func difference(set1 []string, set2 []string) []string {
	var new_str []string
	for _, v := range set1 {
		if !include(v, set2) {
			new_str = append(new_str, v)
		}
	}
	return new_str
}

func main() {
	//文字bi-gramを求める
	//重複したbi-gramは取り除く
	setX := (createNgram("paraparaparadise"))
	setY := (createNgram("paragraph"))

	//XとYの和集合
	fmt.Println(union(setX, setY))

	//XとYの積集合
	fmt.Println(intersection(setX, setY))

	//XとYの差集合
	fmt.Println(difference(setX, setY))

	//"se"という文字列がXとYに含まれるかどうか
	fmt.Println(include("se", setX))
	fmt.Println(include("se", setY))
}
