package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	var word1, word2 string
	fmt.Println("Masukkan kata pertama:")
	scanner.Scan()
	word1 = scanner.Text()

	fmt.Println("Masukkan kata kedua:")
	scanner.Scan()
	word2 = scanner.Text()

	if !isValidLength(word1) || !isValidLength(word2) {
		panic("Input tidak valid: panjang kata tidak boleh lebih dari 10 karakter")
	}

	if !isValidCharacter(word1) || !isValidCharacter(word2) {
		panic("Input tidak valid: hanya diperbolehkan huruf tanpa angka, spasi dan simbol lainnya")
	}
	if areAnagrams(word1, word2) {
		fmt.Printf("%s & %s merupakan anagram\n", word1, word2)
	} else {
		fmt.Printf("%s & %s bukan merupakan anagram\n", word1, word2)
	}
}

func areAnagrams(word1, word2 string) bool {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	if len(word1) != len(word2) {
		return false
	}

	word1Chars := sortString(word1)
	word2Chars := sortString(word2)

	return word1Chars == word2Chars
}

func isValidLength(word string) bool {
	return len(word) <= 10
}

func isValidCharacter(word string) bool {
	for _, char := range word {
		if !isLetter(char) {
			return false
		}
	}
	return true
}

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}
