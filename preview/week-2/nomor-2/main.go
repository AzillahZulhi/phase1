package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

func main() {
	inputFile, err := os.Create("input.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer inputFile.Close()

	data := [][]string{
		{"Name", "Age", "Occupation"},
		{"John", "30", "Engineer"},
		{"Alice", "25", "Doctor"},
		{"Bob", "28", "Teacher"},
		{"Azillah", "24", "Chairman of Parliament"},
	}

	writer := csv.NewWriter(inputFile)
	for _, record := range data {
		if err := writer.Write(record); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()

	fmt.Println("File input berhasil dibuat")

	openInputFile, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer openInputFile.Close()

	reader := csv.NewReader(openInputFile)

	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	outputFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outputFile.Close()

	writer = csv.NewWriter(outputFile)
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var wg sync.WaitGroup

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		wg.Add(1)
		go func(record []string) {
			defer wg.Done()
			processedRecord := processRecord(record)
			err := writer.Write(processedRecord)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}(record)
	}

	wg.Wait()

	writer.Flush()

	fmt.Println("Proses berhasil. Output ditulis di output.csv")

	openOutputFile, err := os.Open("output.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer openOutputFile.Close()

	outputReader := csv.NewReader(openOutputFile)

	for {
		record, err := outputReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println(record)
	}
}

func processRecord(record []string) []string {
	record[0] = capitalize(record[0])
	return record
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	var index int
	for i, ch := range s {
		if unicode.IsLetter(ch) {
			index = i
			break
		}
	}

	return strings.ToUpper(string(s[index])) + s[index+1:]
}
