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
		//正規表現を用いて"Category:"以降の文字列を取得する

		r := regexp.MustCompile(`Category:`)
		if r.MatchString(s) {

			r := regexp.MustCompile(`:(.*)]]`)

			str := r.FindStringSubmatch(s)
			fmt.Println(str[1])
		}
	}

}
