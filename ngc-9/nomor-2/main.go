package main

import (
	"fmt"
	"math"
	"sync"
)

type Circle struct {
	Diameter  float32
	Area      float32
	Perimeter float32
}

func calculateCircleProperties(inputCh <-chan float32, outputCh chan<- Circle, wg *sync.WaitGroup) {
	defer close(outputCh)
	defer wg.Done()

	for diameter := range inputCh {
		radius := diameter / 2
		area := float32(math.Pi * math.Pow(float64(radius), 2))
		perimeter := float32(math.Pi * diameter)
		circle := Circle{
			Diameter:  diameter,
			Area:      area,
			Perimeter: perimeter,
		}
		outputCh <- circle
	}
}

func main() {
	var InputDiameters float32
	fmt.Println("Input Circle Diameter: ")
	fmt.Scanln(&InputDiameters)

	inputCh := make(chan float32)
	outputCh := make(chan Circle)
	var wg sync.WaitGroup

	wg.Add(1)
	go calculateCircleProperties(inputCh, outputCh, &wg)

	wg.Add(1)
	go func() {
		defer close(inputCh)
		defer wg.Done()

		circle := <-outputCh
		fmt.Printf("Diameter: %.2f, Area: %.2f, Perimeter: %.2f\n", circle.Diameter, circle.Area, circle.Perimeter)
	}()

	inputCh <- InputDiameters
	wg.Wait()
}
