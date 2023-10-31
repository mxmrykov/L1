package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	sum := 0
	for _, v := range array {
		wg.Add(1)
		go squareAndPow(v, &sum, &wg)
	}
	wg.Wait()
	fmt.Print(sum)
}

func squareAndPow(num int, sum *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*sum += num * num
}
