package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func scan(sent io.Reader) string {
	scanner := bufio.NewReader(sent)
	input, _ := scanner.ReadString('\n')
	input = strings.Trim(input, "\n")
	return input
}

func main() {
	fmt.Printf("x:")
	x := scan(os.Stdin)

	fmt.Printf("y:")
	y := scan(os.Stdin)

	fmt.Printf("z:")
	z := scan(os.Stdin)

	fmt.Printf("%vの時%vは%v", x, y, z)
}
