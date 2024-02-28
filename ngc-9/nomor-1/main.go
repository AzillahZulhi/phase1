package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGenerator(ch chan<- int) {
	defer close(ch)
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(101)
	ch <- randomNumber
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= 1
	}
	return n * factorial(n-1)
}
func main() {
	randNumCh := make(chan int)
	go randomNumberGenerator(randNumCh)
	var receivedRandomNumber bool

	for num := range randNumCh {
		fmt.Printf("The random number is: %d\n", num)
		fact := factorial(num)
		fmt.Printf("%d! is: %d\n", num, fact)
		receivedRandomNumber = true
		break
	}
	if receivedRandomNumber {
		return
	}
}
