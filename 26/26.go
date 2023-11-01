package main

import (
	"fmt"
	"strings"
)

func main() {
	mapa := map[string]bool{}
	str := "abCdefAf"
	stringa := strings.ToLower(str)

	for i := range stringa {
		if !mapa[string(stringa[i])] {
			mapa[string(stringa[i])] = true
		} else {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
}
