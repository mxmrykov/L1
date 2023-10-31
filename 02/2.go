package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	var wg = sync.WaitGroup{}
	for _, v := range array {
		wg.Add(1)
		go increaseAndPrint(float64(v), &wg)
	}
	wg.Wait()
}

func increaseAndPrint(num float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d: %f | ", int(num), math.Pow(float64(num), 2))
}
