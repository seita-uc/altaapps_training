package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type ApiResponse struct {
	Continue string `json:"continue"`
	Query    struct {
		Pages struct {
			Num struct {
				NS              int    `json:"ns"`
				Title           string `json:"title"`
				Missing         string `json:"missing"`
				Unknown         string `json:"known"`
				Imagerepository string `json:"imagerepository"`
				Imageinfo       []struct {
					Url                 string `json:"url"`
					Descriptionurl      string `json:"descriptionurl"`
					Descriptionshorturl string `json:"descriptionshorturl"`
				} `json:"imageinfo"`
			} `json:"-1"`
		} `json:"pages"`
	} `json:"query"`
}

func template(readFile string) map[string]string {

	//Read File
	inputFile, err := os.Open(readFile)
	f := bufio.NewReader(inputFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	var (
		start bool = false
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
	return dict
}

func main() {
	//テンプレートから国旗画像の情報を取得する
	template := template("formatted.txt")
	imageFile := template["国旗画像"]

	//urlを生成する
	api := "https://www.mediawiki.org/w/api.php?"
	url := fmt.Sprintf(api+"action=query"+"&titles=File:%v"+"&format=json"+"&prop=imageinfo"+"&&iiprop=url", imageFile)
	//スペースを"%20"に置換
	url = strings.Replace(url, " ", "%20", -1)

	res, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Failed to connect")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		fmt.Errorf("Failed to read")
	}

	//	fmt.Printf("%v", string(body))
	var dec ApiResponse
	_ = json.Unmarshal(body, &dec)

	//国旗画像のURL取得
	fmt.Printf("\n%v", dec.Query.Pages.Num.Imageinfo[0].Url)
}
