package main

import (
	"fmt"
)

func BubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func main() {
	a := []int{4, 46, 77, 89, 94, 32, 33, 11, 45, 67, 89, 7, 64, 33, 67, 85, 89, 0, 22, 27, 31}
	fmt.Println(BubbleSort(a))
}
