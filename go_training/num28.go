package main

import (
	"bufio"
	"flag"
	"fmt"
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
	markUp1 := regexp.MustCompile(`(.*)<ref.+>`)
	markUp2 := regexp.MustCompile(`\{\{.*[Ll]ang\|[a-zA-Z\-]+\|(.+)\}\}`)
	markUp3 := regexp.MustCompile(`'+(.+?)'+`)
	markUp4 := regexp.MustCompile(`(.*?)<br\s?/?>`)
	http := regexp.MustCompile(`http:.*`)
	internalLink := regexp.MustCompile(`[\[\]]`)

	for {
		s, _ := f.ReadString('\n')
		if start {
			if contents.MatchString(s) {
				tmpStr := internalLink.ReplaceAllString(s, "")
				tmpStr = http.ReplaceAllString(tmpStr, "")

				if markUp1.MatchString(tmpStr) {
					m := markUp1.FindStringSubmatch(tmpStr)
					tmpStr = markUp1.ReplaceAllString(tmpStr, m[1])
				}

				if markUp2.MatchString(tmpStr) {
					m := markUp2.FindStringSubmatch(tmpStr)
					tmpStr = markUp2.ReplaceAllString(tmpStr, m[1])
				}

				if markUp3.MatchString(tmpStr) {
					m := markUp3.FindStringSubmatch(tmpStr)
					tmpStr = markUp3.ReplaceAllString(tmpStr, m[1])
				}

				if markUp4.MatchString(tmpStr) {
					m := markUp4.FindStringSubmatch(tmpStr)
					tmpStr = markUp4.ReplaceAllString(tmpStr, m[1])
				}
				infoStr := contents.FindStringSubmatch(tmpStr)
				dict[infoStr[1]] = infoStr[2]
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
