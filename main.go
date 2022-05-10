package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsPalindrome_1("f"))
}

func IsPalindrome_1(s string) bool {
	for i := 0; i < len(s); i++ {
		j := len(s) - 1 - i
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func IsPalindrome_2(s string) bool {
	res := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}

	return s == string(res)
}
