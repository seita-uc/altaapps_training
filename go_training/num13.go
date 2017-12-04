package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	sourceFile, _ := os.Open(`hoge.txt`)
	file1, _ := os.Create("./col1.txt")

	file2, _ := os.Create("./col2.txt")

	scanner := bufio.NewScanner(sourceFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		newString := strings.Split(s, "\t")
		file1.WriteString(newString[0] + "\n")
		file2.WriteString(newString[1] + "\n")
	}

}
