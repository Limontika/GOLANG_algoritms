package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 10, 101, 2, 9, 20, 100}
	for i := 0; i < len(a)-1; i++ {
		if a[i+1] < a[i] {
			for j := i + 1; j <= i; j-- {
				if a[i] > a[j] {
					a[i], a[j] = a[j], a[i]
				}
			}
		}
	}
	fmt.Println(a)
}
