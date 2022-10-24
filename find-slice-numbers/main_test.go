package main

import "testing"

func TestFindMin(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{
			data:   []int{100, 2, 44, 55, 4},
			answer: 2,
		},
		{
			data:   []int{800, 1, 55, 77, 100, 20000, 2121212},
			answer: 1,
		}, {
			data:   []int{600, 33333, 12123123, 12312344, 44555553},
			answer: 600,
		},
	}
	for _, v := range tests {
		x := findMin(v.data)
		if x != v.answer {
			t.Error("Got", v.answer, "Expected", x)
		}
	}
}

func TestFindMax(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{
			data:   []int{100, 2, 44, 55, 4},
			answer: 100,
		},
		{
			data:   []int{800, 1, 55, 77, 100, 20000, 2121212},
			answer: 2121212,
		}, {
			data:   []int{600, 33333, 12123123, 12312344, 44555553},
			answer: 44555553,
		},
	}
	for _, v := range tests {
		x := findMax(v.data)
		if x != v.answer {
			t.Error("Got", v.answer, "Expected", x)
		}
	}
}
