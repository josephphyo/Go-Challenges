package main

import "fmt"

func partition(x []int, low, high int) ([]int, int) {
	pivot := x[high]
	i := low
	for j := low; j < high; j++ {
		if x[j] < pivot {
			x[i], x[j] = x[j], x[i]
			i++
		}
	}
	x[i], x[high] = x[high], x[i]
	return x, i
}

func quickSort(x []int, low, high int) []int {
	if low < high {
		x, p := partition(x, low, high)
		x = quickSort(x, low, p-1)
		x = quickSort(x, p+1, high)
	}
	return x
}

func main() {
	xi := []int{10, 6, 100, 20, 50, 30, 2, 1, 5,
		8, 3, 4, 7, 9}
	fmt.Println("Quick Sort --> ", quickSort(xi, 0, len(xi)-1))
}
