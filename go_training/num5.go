package main

import "fmt"

func main() {
	str := []
	numbers := 0
	str = fmt.Sprintf("%v", numbers)
	str = strings.Trim(str, "[]")
	str = strings.Replace(str, " ", ",", -1)
}
