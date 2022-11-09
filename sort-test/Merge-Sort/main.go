package main

import "fmt"

func split(xi []int) []int {
	if len(xi) < 2 {
		return xi
	}
	mid := len(xi) / 2
	//fmt.Println(mid)
	leftArr := xi[:mid]
	rightArr := xi[mid:]
	//fmt.Println("Left Array ->", leftArr)
	//fmt.Println("Right Array ->", rightArr)
	return mergeSort(split(leftArr), split(rightArr))
}

func mergeSort(leftArr, rightArr []int) []int {
	results := []int{}
	lIndex := 0
	rIndex := 0

	for lIndex < len(leftArr) && rIndex < len(rightArr) {
		if leftArr[lIndex] > rightArr[rIndex] {
			results = append(results, rightArr[rIndex])
			rIndex++
		} else {
			results = append(results, leftArr[lIndex])
			lIndex++
		}
	}
	for ; lIndex < len(leftArr); lIndex++ {
		fmt.Println("The Lef\"The Left Array ->\"t Array ->", leftArr, "The Right Array ->", rightArr)
		results = append(results, leftArr[lIndex])
	}
	for ; rIndex < len(rightArr); rIndex++ {
		fmt.Println("The Right Array ->", rightArr, "The Left Array ->", leftArr)
		results = append(results, rightArr[rIndex])
	}
	return results
}

func main() {
	xi := []int{10, 6, 100, 20, 50, 30, 2, 1, 5,
		8, 3, 4, 7, 9}
	fmt.Println(split(xi))
}
