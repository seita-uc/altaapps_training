package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func countLines(sourceFile *os.File) int {
	scanner := bufio.NewScanner(sourceFile)
	scanner.Split(bufio.ScanLines)
	num := 0

	for scanner.Scan() {
		num++
	}
	return num
}

func necessaryNumbers(arg int, lineNum int) (leftNum, splitNum int) {
	leftNum = lineNum % arg
	splitNum = (lineNum - leftNum) / arg
	return leftNum, splitNum
}

func scan(num int, readFile string) string {

	sourceFile, _ := os.Open(readFile)
	scanner := bufio.NewScanner(sourceFile)
	for i := 1; scanner.Scan(); i++ {
		if num == i {

			return scanner.Text()
		}

	}
	return scanner.Text()
}

func main() {
	//フラグ処理
	var (
		arg      int
		readFile string
	)

	//go run num17.go -s 分割回数 -f ファイル名
	flag.IntVar(&arg, "s", 1, "分割回数")
	flag.StringVar(&readFile, "f", "", "読み込みファイル")
	flag.Parse()

	sourceFile, _ := os.Open(readFile)
	lineNum := countLines(sourceFile)
	leftNum, splitNum := necessaryNumbers(arg, lineNum)

	n := 1

	for i := 1; i <= splitNum+1; i++ {

		filename := fmt.Sprintf("./split%d.txt", i)
		for file, _ := os.Create(filename); ; n++ {
			s := scan(n, readFile)

			if n > arg*splitNum {
				for x := 0; x <= leftNum; x++ {
					s := scan(n+x, readFile)

					file.WriteString(s + "\n")
				}
				break

			}

			if n <= arg*i {
				file.WriteString(s + "\n")

			}
			if n == arg*i {
				n++
				break
			}
		}
	}
}
