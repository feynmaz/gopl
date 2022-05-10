package main

import "testing"

func BenchmarkIsPalindrome_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome_1("A man,  a plan,  a canal:  Panama")
	}
}

func BenchmarkIsPalindrome_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome_2("A man,  a plan,  a canal:  Panama")
	}
}
