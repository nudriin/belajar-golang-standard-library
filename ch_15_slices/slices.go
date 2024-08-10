package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Nurdin", "NestJS", "React", "ExpressJS", "Golang"}
	values := []int{10, 200, 312, 31, 22, 441, 565, 13, 87, 3, 54, 75, 56}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(names, "NestJS"))
	fmt.Println(slices.Contains(names, "Joseph"))
	fmt.Println(slices.Contains(values, 31))
	fmt.Println(slices.Index(values, 31))
	fmt.Println(slices.Index(names, "NestJS"))
	fmt.Println(slices.Index(names, "Aras"))
}
