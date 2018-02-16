package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//フラグ処理
	var (
		readFile string
		key      int
	)

	//go run num18.go -f ファイル名 -k キーの数
	flag.StringVar(&readFile, "f", "", "読み込みファイル")
	flag.IntVar(&key, "k", 1, "ソート対象になるキーの数")
	flag.Parse()

	//read file
	input_file, err := os.Open(readFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	sMap := make(map[float64]string)
	var newString []string
	scanner := bufio.NewScanner(input_file)

	for scanner.Scan() {
		s := scanner.Text()
		tmpString := strings.Fields(s)
		newString = append(newString, tmpString[key])
		i, _ := strconv.ParseFloat(tmpString[key], 64)
		sMap[i] = s
	}

	newInts := []float64{}

	for _, v := range newString {
		i, _ := strconv.ParseFloat(v, 64)
		newInts = append(newInts, i)
	}

	sort.Float64s(newInts)
	for i := len(newInts) - 1; i >= 0; i-- {
		fmt.Println(sMap[newInts[i]])
	}

}
