package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	temps := [4]float64{28.8, 10.4, -34.3, 22.2}
	readyTemp := map[int][]float64{}

	for _, temp := range temps {
		readyTemp[int(temp/10)*10] = append(readyTemp[int(temp/10)*10], temp)
	}

	arr, err := json.MarshalIndent(readyTemp, "", "\t")

	if err != nil {

	}
	fmt.Println(string(arr))
}
