package main

import (
	"fmt"
	"time"
)

func FizzBuzz(channel1 <-chan []int, outputChannel1 chan<- int) {
	total, ganjil, genap := 0, 0, 0
	for _, value := range <-channel1 {
		if value%3 == 0 && value%5 == 0 {
			fmt.Println(value, "FizzBuzz")
		} else if value%3 == 0 {
			fmt.Println(value, "Fizz")
		} else if value%5 == 0 {
			fmt.Println(value, "Buzz")
		} else {
			fmt.Println(value)
		}
		if value%2 == 0 {
			genap += value
		} else {
			ganjil += value
		}
		total += value
	}
	outputChannel1 <- genap
	outputChannel1 <- ganjil
	fmt.Println("Total:", total)
	time.Sleep(2 * time.Second)
}

func main() {
	channel1 := make(chan []int)
	outputChannel1 := make(chan int, 2)
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	go func() {
		channel1 <- a
	}()
	go FizzBuzz(channel1, outputChannel1)
	time.Sleep(3 * time.Second)
}
