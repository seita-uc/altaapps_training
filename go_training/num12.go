package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sourceFile, _ := os.Open(`hoge.txt`)
	scanner := bufio.NewScanner(sourceFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		str := scanner.Text()
		resStr := strings.Replace(str, "\t", " ", -1)
		fmt.Println(resStr)
	}

}
