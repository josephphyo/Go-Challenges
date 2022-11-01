package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	var acount int
	var bcount int
	var result []int32

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			acount++
		} else if a[i] < b[i] {
			bcount++
		}
	}
	result = append(result, int32(acount), int32(bcount))
	return result
}

func main() {
	a := []int32{13, 46, 7}
	b := []int32{3, 6, 10}

	fmt.Println(compareTriplets(a, b))
}
