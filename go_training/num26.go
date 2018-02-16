package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		readFile string
	)

	flag.StringVar(&readFile, "f", "", "read file")
	flag.Parse()

	//Read File
	inputFile, err := os.Open(readFile)
	f := bufio.NewReader(inputFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	var (
		start bool = false
		//		field string
		//		value string
	)

	dict := make(map[string]string)

	startInfo := regexp.MustCompile(`^\{\{基礎情報.*`)
	contents := regexp.MustCompile(`^\|(.*)\s=\s(.*)`)
	endInfo := regexp.MustCompile(`^\}\}\n`)
	tag := regexp.MustCompile(`<.*?>`)
	http := regexp.MustCompile(`http.*`)

	for {
		s, _ := f.ReadString('\n')
		if start {
			if contents.MatchString(s) {
				infoStr := contents.FindStringSubmatch(s)
				//			tmpStr := strings.Replace(infoStr[2], "[", "", -1)
				//			tmpStr = strings.Replace(tmpStr, "]", "", -1)
				tmpStr := strings.Replace(infoStr[2], "'", "", -1)

				if tag.MatchString(tmpStr) {
					tmpStr = tag.ReplaceAllString(tmpStr, "")
				}

				if http.MatchString(tmpStr) {
					tmpStr = http.ReplaceAllString(tmpStr, "")
				}

				dict[infoStr[1]] = tmpStr
			}
		}
		if startInfo.MatchString(s) {
			start = true
		}
		if endInfo.MatchString(s) {
			break
		}

	}

	for k, v := range dict {
		fmt.Printf("(%v)	%v\n", k, v)
	}
}
