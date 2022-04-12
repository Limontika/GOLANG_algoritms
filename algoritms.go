package main

import (
	"fmt"
)

func main() {
	input := []int{100, 80, 30, 35, 20, 10, 5, 2, 1, 0}
	buble := bubleSort(input)
	insert := insertSort(input)
	fmt.Println(buble, insert)
}

func bubleSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}

func insertSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] < input[i] {
			for j := i + 1; j <= i; j-- {
				if input[i] > input[j] {
					input[i], input[j] = input[j], input[i]
				}
			}
		}
	}
	return input
}
