package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

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
		r := regexp.MustCompile(`(?:File|ファイル):(.+?)\|`)
		newStr := r.FindStringSubmatch(s)
		if newStr != nil {

			fmt.Printf("%v \n", newStr[1])

		}

	}

}