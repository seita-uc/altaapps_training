package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func levelCount(line string, s string) int {
	level := 0
	for _, v := range line {
		if string(v) == s {
			level++
		}
	}
	return level/2 - 1
}

func main() {
	var (
		readFile string
	)
	flag.StringVar(&readFile, "f", "", "read file")
	flag.Parse()

	//Read File
	inputFile, err := os.Open(readFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	file := bufio.NewReader(inputFile)

	for {
		s, err := file.ReadString('\n')
		if err == io.EOF {
			break
		}
		r := regexp.MustCompile(`^=.*=`)
		if r.MatchString(s) {
			level := levelCount(s, "=")
			s = strings.Trim(s, "=\n")
			fmt.Printf("%v (%v)\n", s, level)
		}
	}

}
