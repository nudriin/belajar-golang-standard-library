package main

import (
	"fmt"
	"strconv"
)

func main() {
	// parsing tipe data lain ke string

	fmt.Println(strconv.ParseBool("true")) // mengubah string menjadi boolean
	fmt.Println(strconv.Atoi("20"))        // mengubah string menjadi Int

	// Mengubah decimal ke binary
	binnary := strconv.FormatInt(10, 2)
	fmt.Println(binnary)

	// Mengubah boolean ke string
	boolString := strconv.FormatBool(true)
	fmt.Println(boolString)

	// Mengubah int ke string
	intString := strconv.Itoa(1000)
	fmt.Println(intString)
}
