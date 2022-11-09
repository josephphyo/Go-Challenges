package main

import "fmt"

func PalinDrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("ror is Palindrome?", PalinDrome("ror"))
	fmt.Println("go is Palindrome?", PalinDrome("go"))

}
