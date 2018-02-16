package main

//各行の1コラム目の文字列の出現頻度を求め，出現頻度の高い順に並べる

/*19. 各行の1コラム目の文字列の出現頻度を求め，出現頻度の高い順に並べる
各行の1列目の文字列の出現頻度を求め，その高い順に並べて表示せよ．
確認にはcut, uniq, sortコマンドを用いよ．*/

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type ValSorter struct {
	Keys []string
	Vals []int
}

func NewValSorter(m map[string]int) *ValSorter {
	vs := &ValSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]int, 0, len(m)),
	}
	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Vals = append(vs.Vals, v)
	}
	return vs
}

func (vs *ValSorter) Sort() {
	sort.Sort(vs)
}

func (vs *ValSorter) Len() int { return len(vs.Vals) }

func (vs *ValSorter) Less(i, j int) bool { return vs.Vals[i] > vs.Vals[j] } //降順なので逆
func (vs *ValSorter) Swap(i, j int) {
	vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
	vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
}

func main() {

	//フラグ処理
	var (
		readFile string
	)

	//go run num18.go -f ファイル名
	flag.StringVar(&readFile, "f", "", "読み込みファイル")
	flag.Parse()

	//read file
	input_file, err := os.Open(readFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input_file)

	listOfCol := make(map[string]int)

	//ファイルの各行を見ていく
	for scanner.Scan() {
		str := scanner.Text()
		cols := strings.Fields(str)

		listOfCol[cols[0]] += 1
		fmt.Println(listOfCol)
		if err == io.EOF {
			break
		}

	}

	vs := NewValSorter(listOfCol)
	vs.Sort()

	for i := range vs.Keys {
		fmt.Printf("%v %v\n", vs.Keys[i], vs.Vals[i])
	}
}
