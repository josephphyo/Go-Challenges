package main

import (
	"fmt"
	"testing"
)

func TestPalinDrome(t *testing.T) {
	type test struct {
		data   []string
		answer bool
	}

	tests := []test{
		{
			data:   []string{"ror", "gog", "oko"},
			answer: true,
		},
		{
			data:   []string{"fo", "ok", "no", "jo", "fd"},
			answer: false,
		},
	}
	for _, v := range tests {
		for _, j := range v.data {
			//fmt.Printf("%v\n", j)
			p := PalinDrome(j)
			fmt.Println("input->", j, "result->", p)
			if p != v.answer {
				t.Error("Got", v.answer, "Expected", p)
			}
		}
	}
}
