package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	num := 0
	file, _ := os.Open(`hoge.txt`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num++
	}
	fmt.Println(num)
}
