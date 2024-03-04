package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("{error}")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("Hello dunia nama saya Azillah Zulhi\n")
	if err != nil {
		fmt.Println("{error}")
		return
	}

	_, err = writer.WriteString("Hello dunia nama saya nama nama Hello hello\n")
	if err != nil {
		fmt.Println("{error}")
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("{error}")
		return
	}

	fmt.Println("File berhasil dibuat")

	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("{error}")
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			wg.Add(1)
			go func(w string) {
				defer wg.Done()
				mutex.Lock()
				wordCount[w]++
				mutex.Unlock()
			}(word)
		}
	}
	wg.Wait()

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
