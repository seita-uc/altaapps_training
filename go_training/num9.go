package main

//与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ．
//
//英小文字ならば(219 - 文字コード)の文字に置換
//その他の文字はそのまま出力
//この関数を用い，英語のメッセージを暗号化・復号化せよ．

import (
	"fmt"
	"regexp"
)

func check_regexp(reg, str string) bool {
	r := regexp.MustCompile(reg).MatchString(str)
	return r
}

func cipher(s string) string {

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {

		if check_regexp(`[a-z]`, string(runes[i])) {
			runes[i] = 219 - runes[i]
		}
	}
	return string(runes)
}

func main() {
	fmt.Printf(cipher("I wanna sleep in the bed"))

}
