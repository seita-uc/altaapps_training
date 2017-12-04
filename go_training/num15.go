package main

// 自然数Nをコマンドライン引数などの手段で受け取り，入力のうち先頭のN行だけを表示せよ．確認にはheadコマンドを用いよ．

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func printLines(arg int) {

	sourceFile, _ := os.Open(`hoge.txt`)
	scanner := bufio.NewScanner(sourceFile)
	scanner.Split(bufio.ScanLines)
	num := 0

	for scanner.Scan() {
		if num < arg {
			line := scanner.Text()
			fmt.Println(line)
			num++
		}
	}
}

func main() {

	//フラグ処理
	var arg int
	flag.IntVar(&arg, "N", 1, "表示する行数")
	flag.Parse()

	fmt.Println("argument:", arg)
	printLines(arg)

}
