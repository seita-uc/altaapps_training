package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type TextTitle struct {
	Text  string `json:"text"`
	Title string `json:"title"`
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
	writefile, _ := os.Create("./formatted.txt")

	for {
		s, err := file.ReadString('\n')
		if err == io.EOF {
			break
		}

		byteText := []byte(s)
		var tt TextTitle
		_ = json.Unmarshal(byteText, &tt)

		if tt.Title == "イギリス" {

			fmt.Printf("%v\n", tt.Text)
			writefile.WriteString(tt.Text + "\n")

		}
	}
}
