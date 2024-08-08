package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// membuat ring dengan 5 data
	data := ring.New(5)

	// mengisi data pada ring
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// cara mengakses valuenya
	data.Do(func(data any) {
		fmt.Println(data)
	})

}
