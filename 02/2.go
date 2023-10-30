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
		go func(v int) {
			defer wg.Done()
			fmt.Printf("%d: %f | ", v, math.Pow(float64(v), 2))
		}(v)
	}
	wg.Wait()
}
