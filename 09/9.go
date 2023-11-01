package main

import (
	"fmt"
)

func main() {
	array := [5]int{0, 1, 2, 3, 4}

	ch1 := make(chan int)
	ch2 := make(chan int)
	chFlag := make(chan bool)

	go func() {
		for {
			select {
			case rm := <-ch1:
				ch2 <- rm * 2
			case <-chFlag:
				close(ch1)
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case r := <-ch2:
				fmt.Printf("%d ", r)
			case <-chFlag:
				fmt.Printf("%d ", <-ch2)
				close(ch2)
				return
			}
		}
	}()

	for i := 0; i < len(array); i += 1 {
		ch1 <- array[i]
	}
	chFlag <- true
	close(chFlag)
}
