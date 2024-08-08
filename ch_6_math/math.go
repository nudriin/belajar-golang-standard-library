package main

import (
	"fmt"
	"math"
)

func main() {
	// membulatkan ke atas atau ke bawah
	round := math.Round(10.6)
	fmt.Println(round)

	// membulatkan ke atas
	ceil := math.Ceil(10.2)
	fmt.Println(ceil)

	// membulatkan ke atas
	floor := math.Floor(10.6)
	fmt.Println(floor)

	// mencari nilai tertinggi
	max := math.Max(10.6, 10.1)
	fmt.Println(max)

	// mencari nilai terendah
	min := math.Min(10.6, 10.1)
	fmt.Println(min)

}
