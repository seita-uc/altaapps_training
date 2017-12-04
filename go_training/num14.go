package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	sourceFile1, _ := os.Open(`col1.txt`)
	sourceFile2, _ := os.Open(`col2.txt`)
	file, _ := os.Create("./foo.txt")

	scanner1 := bufio.NewScanner(sourceFile1)
	scanner1.Split(bufio.ScanWords)

	scanner2 := bufio.NewScanner(sourceFile2)
	scanner2.Split(bufio.ScanWords)

	for scanner1.Scan() && scanner2.Scan() {
		col1 := scanner1.Text()
		col2 := scanner2.Text()
		newString := []string{col1, col2}
		result := strings.Join(newString, "\t")
		file.WriteString(result + "\n")
	}

}
