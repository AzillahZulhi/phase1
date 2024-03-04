package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
)

func main() {

	cpuFile, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer cpuFile.Close()
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	var num1, num2 float64
	fmt.Println("Masukkan angka pertama:")
	if scanner.Scan() {
		input := scanner.Text()
		var err error
		num1, err = strconv.ParseFloat(input, 64)
		if err != nil {
			panic("Input tidak valid: hanya diperbolehkan angka saja")
		}
	}

	fmt.Println("Masukkan angka kedua:")
	if scanner.Scan() {
		input := scanner.Text()
		var err error
		num2, err = strconv.ParseFloat(input, 64)
		if err != nil {
			panic("Input tidak valid: hanya diperbolehkan angka saja")
		}
	}
	fmt.Printf("Angka pertama: %.2f\n", num1)
	fmt.Printf("Angka kedua: %.2f\n", num2)

	plusResult := plus(num1, num2)
	minusResult := minus(num1, num2)
	multiplicationResult := multiplication(num1, num2)
	divisionResult, err := division(num1, num2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Hasil penambahan %.2f dan %.2f adalah: %.2f\n", num1, num2, plusResult)
	fmt.Printf("Hasil pengurangan %.2f dan %.2f adalah: %.2f\n", num1, num2, minusResult)
	fmt.Printf("Hasil perkalian %.2f dan %.2f adalah: %.2f\n", num1, num2, multiplicationResult)
	fmt.Printf("Hasil pembagian %.2f dan %.2f adalah: %.2f\n", num1, num2, divisionResult)

	memFile, err := os.Create("mem.pprof")
	if err != nil {
		fmt.Println("Error while creating memory file", err)
	}
	defer memFile.Close()

	runtime.GC()
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Println("Error starting CPU Profile", err)
		return
	}
}

func plus(a, b float64) float64 {
	return a + b
}
func minus(a, b float64) float64 {
	return a - b
}
func multiplication(a, b float64) float64 {
	return a * b
}
func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Pembagian oleh nol")
	}
	return a / b, nil
}
