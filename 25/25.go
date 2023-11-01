package main

import (
	"fmt"
	"time"
)

func main() {
	timeToSleep := 5
	fmt.Println("Prog start")
	Sleep(time.Duration(timeToSleep))
	fmt.Println("Prog end")
}

func Sleep(duration time.Duration) {
	<-time.After(duration * time.Second)
}
