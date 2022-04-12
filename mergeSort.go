package main

import "fmt"

func main() {
	input := []int{80, 100, 30, 50, 9, 20, 85, 5, 2, 1, 0}
	sortMerge := mergeSort(input)
	fmt.Println(sortMerge)
}

func mergeSort(input []int) []int {
	if len(input) == 1 {
		return input
	}
	n := len(input)
	a := input[:n/2]
	b := input[n/2:]
	left := mergeSort(a)
	right := mergeSort(b)
	return merge(left, right)
}

func merge(a []int, b []int) []int {
	var output []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			output = append(output, a[i])
			i++
		} else {
			output = append(output, b[j])
			j++
		}
	}
	for i < len(a) {
		output = append(output, a[i])
		i++
	}
	for j < len(b) {
		output = append(output, b[j])
		j++
	}
	return output
}
