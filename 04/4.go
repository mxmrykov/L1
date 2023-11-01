package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	mainChanel := make(chan int)
	const workerCount = 5
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i += 1 {
		wg.Add(1)
		go worker(mainChanel, &wg, i)
	}
	info := make(chan os.Signal, 1)
	signal.Notify(info, syscall.SIGINT, os.Interrupt)
	for {
		select {
		case <-info:
			close(mainChanel)
			close(info)
			return
		default:
			mainChanel <- rand.Intn(100)
			time.Sleep(2 * time.Second)
		}
	}
	wg.Wait()
}

func worker(ch chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("mes: %d; WID: %d\n", msg, workerId)
	}
}
