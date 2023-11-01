package main

import "fmt"

func setBit(n, i int64) int64 {
	n = n << (i - 1)
	n = n | 1
	return n >> (i - 1)
}

func unsetBit(n, i int64) int64 {
	return 0
}

func main() {
	a := int64(18)
	fmt.Println(a)
	a = setBit(a, 5)
	fmt.Println(a)
}
