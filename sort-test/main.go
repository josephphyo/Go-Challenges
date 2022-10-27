package main

import "fmt"

func main() {
	xi := []int{10, 6, 100, 20, 50, 30, 2, 1, 5,
		8, 3, 4, 7, 9}

	for i := 0; i < len(xi); i++ {
		for j := 0; j < len(xi); j++ {
			if xi[j] > xi[i] {
				xi[j], xi[i] = xi[i], xi[j]
			}
		}
	}
	fmt.Println("Selection Sort", xi)
}
