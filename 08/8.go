package main

import "fmt"

func setBit(n, i int64) int64 {
	return n | (1 << (i - 1))
}

func unsetBit(n, i int64) int64 {
	return n &^ (1 << (i - 1))
}

func main() {
	a := int64(18)
	fmt.Println(a)
	a = setBit(a, 2)
	fmt.Println(a)
	a = unsetBit(a, 2)
	fmt.Println(a)
}
