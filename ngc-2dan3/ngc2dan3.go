package main

import (
	"fmt"
	// "sort"
	// "golang.org/x/tools/go/analysis/passes/sortslice"
	// "math/rand"
	// "time"
)

func main() {
	// conditonal1()
	// conditonal2()
	// conditonal3()
	// conditonal4()
	// conditonal5()
	// conditonal6()
	// conditonal7()
	// conditonal8()
	conditonal9()
}

// ngc 2
// func conditonal1() {

// 	rand.Seed(time.Now().UnixNano())

// 	var name string
// 	fmt.Println("Masukkan nama Anda: ")
// 	fmt.Scanln(&name)

// 	randomNumber := rand.Intn(100) + 1

// 	switch {
// 	case randomNumber > 80:
// 		fmt.Printf("Selamat %v, anda sangat beruntung\n", name)
// 	case randomNumber <= 80 && randomNumber > 60:
// 		fmt.Printf("Selamat %v, anda beruntung\n", name)
// 	case randomNumber <= 60 && randomNumber > 40:
// 		fmt.Printf("Mohon maaf %v, anda sangat kurang beruntung\n", name)
// 	default:
// 		fmt.Printf("Mohon maaf %v, anda sial\n", name)
// 	}
// }

// func conditonal2() {
// 	var age int
// 	fmt.Println("Masukkan umur Anda: ")
// 	fmt.Scanln(&age)

// 	if age < 0 || age > 100 {
// 		fmt.Println("Input tidak valid")
// 	} else if age >= 18 {
// 		fmt.Println("Silahkan masuk")
// 	} else {
// 		fmt.Println("Dilarang masuk, buat KTP dulu sono")
// 	}
// }

//ngc 3

// func conditonal3() {
// 	var names = []string{"Hank", "Heisenberg", "Skyler"}
// 	var ages = []int{50, 52, 48}
// 	var jobs = []string{"Polisi", "Ilmuan", "Akuntan"}

// 	for i := 0; i < len(names); i++ {
// 		fmt.Printf("Hallo nama saya %s, umur saya %d, pekerjaan saya %s\n", names[i], ages[i], jobs[i])
// 	}
// }

// func conditonal4() {
// 	var num1 = []float64{1, 5, 7, 8, 10, 24, 33}
// 	var num2 = []float64{1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

// 	totalNum1 := 0.0
// 	for _, nilai := range num1 {
// 		totalNum1 += nilai
// 	}
// 	rataNum1 := totalNum1 / float64(len(num1))

// 	totalNum2 := 0.0
// 	for _, nilai := range num2 {
// 		totalNum2 += nilai
// 	}
// 	rataNum2 := totalNum2 / float64(len(num2))

// 	allNumbers := append(num1, num2...)

// 	sort.Float64s(allNumbers)

// 	n := len(allNumbers)

// 	var median float64
// 	if n%2 == 0 {
// 		median = (allNumbers[n/2-1] + allNumbers[n/2]) / 2
// 	} else {
// 		median = allNumbers[n/2]
// 	}

// 	fmt.Printf("Rata-rata data pertama: %.2f\n", rataNum1)
// 	fmt.Printf("Rata-rata data kedua: %.2f\n", rataNum2)
// 	fmt.Printf("Median dari kedua set data: %.2f\n", median)
// }

// func conditonal5() {
// 	numbers := []int{5, 3, 9, 7, 1, 4, 8, 2, 6}
// 	for i := 0; i < len(numbers); i++ {
// 		for j := i + 1; j < len(numbers); j++ {
// 			if numbers[i] < numbers[j] {
// 				temp := numbers[i]
// 				numbers[i] = numbers[j]
// 				numbers[j] = temp
// 			}
// 		}
// 	}
// 	fmt.Println("Slice setelah diurutkan dari besar ke kecil:", numbers)
// }

// func conditonal6() {
// 	rows := 5

// 	for i := 0; i < rows; i++ {
// 		fmt.Println("*")
// 	}
// }

// func conditonal7() {
// 	rows := 5
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < rows; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func conditonal8() {
// 	rows := 5
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j <= i; j++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

func conditonal9() {
	rows := 5
	for i := 0; i < rows; i++ {
		for j := 0; j < rows-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
