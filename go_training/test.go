package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	for _, v := range os.Args {
		fmt.Println(v)
	}

}
