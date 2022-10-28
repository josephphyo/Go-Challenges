package main

import "fmt"

func main() {
	xi := []int{10, 6, 100, 20, 50, 30, 2, 1, 5,
		8, 3, 4, 7, 9}

	//fmt.Println(len(xi))

	for i := 0; i < len(xi); i++ {
		for j := 0; j < len(xi)-1; j++ {
			if xi[j] > xi[j+1] {
				xi[j], xi[j+1] = xi[j+1], xi[j]
			}
		}
	}
	fmt.Println("Bubble Sort:", xi)
}
