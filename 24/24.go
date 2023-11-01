package main

import (
	"fmt"
	"image"
	"math"
)

func main() {
	dot1 := image.Point{X: 1, Y: 1}
	dot2 := image.Point{X: 2, Y: 2}

	fmt.Printf("dist: %.2f", calcDist(&dot1, &dot2))
}

func calcDist(p1, p2 *image.Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}
