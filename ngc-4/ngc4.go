package main

import (
	"fmt"
	"strings"
)

func main() {
	// Menggunakan AlayGen dengan argumen yang benar
	fmt.Println(AlayGen("hello", "world", "zzz"))

	// Menggunakan Fibonacci dengan argumen yang benar
	n := 10
	fib := Fibonacci(n)
	fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", n, fib)
}

func AlayGen(words ...string) string {
	wordsAlay := toAlay(words...)
	result := strings.Join(wordsAlay, " ")
	return result
}

func toAlay(words ...string) []string {
	alayMap := map[rune]string{
		'a': "4",
		'e': "3",
		'i': "!",
		'l': "1",
		'n': "N",
		's': "5",
		'x': "*",
	}
	var result []string
	for _, word := range words {
		var alayWord string

		for _, char := range word {
			if alay, ok := alayMap[char]; ok {
				alayWord += alay
			} else {
				alayWord += string(char)
			}
		}
		result = append(result, alayWord)
	}
	return result
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
