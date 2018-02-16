/*17. １列目の文字列の異なり
1列目の文字列の種類（異なる文字列の集合）を求めよ．確認にはsort, uniqコマンドを用いよ．
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

// n がスライスに含まれているか
func include(n string, xs []string) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

func main() {

	//フラグ処理
	var (
		readFile string
	)

	//go run num18.go -f ファイル名
	flag.StringVar(&readFile, "f", "", "読み込みファイル")
	flag.Parse()

	//?????
	//	argNum := flag.NArg()
	//	if len(argNum) != 2 {
	//		fmt.Fprint(os.Stderr, "not enough args\n")
	//		os.Exit(1)
	//	}

	//input hoge.txt
	input_file, err := os.Open(readFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	file := bufio.NewReader(input_file)

	uniq_strings := []string{}

	for {
		str, err := file.ReadString('\n')
		if err == io.EOF {
			break
		}
		if !include(str, uniq_strings) {
			uniq_strings = append(uniq_strings, str)
		}

	}

	sort.Strings(uniq_strings)

	for i := range uniq_strings {
		fmt.Printf("%v", uniq_strings[i])
	}

}
