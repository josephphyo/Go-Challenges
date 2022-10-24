package main

import "fmt"

func main() {
	array1 := []int{23, 26, 49, 78, 55,
		50, 100, 49, 72, 53, 3, 4,
		94, 45, 89, 72, 43,
		62, 67, 89, 56, 92,
		56, 76, 80, 49, 74,
		61, 69, 30, 88, 99,
		36, 33, 48, 65, 64,
		88, 69, 90, 43, 82,
		67, 25, 93, 36, 23,
		28, 85, 43, 48, 55,
		35, 63, 66, 43, 54,
		32, 94, 69, 56, 62,
		83, 56, 55, 66, 75,
		61, 99, 56, 36, 72}
	array2 := []int{
		30, 29, 18, 35, 86,
		32, 109, 95, 10, 60,
		28, 48, 35, 6, 89,
		79, 41, 8, 60, 75,
		94, 81, 87, 30, 47,
		81, 53, 55, 57, 16,
		50, 47, 78, 100, 93,
		53, 25, 16, 75, 12,
		96, 63, 12, 17, 40,
		92, 13, 9, 32, 64,
		55, 96, 24, 48, 65,
		16, 58, 6, 4, 11,
		91, 28, 73, 72, 30,
		31, 11, 75, 67, 26,
		54, 59, 26, 13, 83,
		18, 94, 100, 26, 47,
		4, 71, 64, 66, 69,
	}

	fmt.Println("Min Number of slice1 : ")
	findMin(array1)
	fmt.Println("Min Number of slice2: ")
	findMin(array2)
	fmt.Println("-----------------")
	fmt.Println("Max Number of slice1 : ")
	findMax(array1)
	fmt.Println("Max Number of slice2: ")
	findMax(array2)
}

func findMin(n []int) int {
	min := n[0]
	for _, v := range n {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
	return min
}

func findMax(n []int) int {
	max := n[0]
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
	return max
}
