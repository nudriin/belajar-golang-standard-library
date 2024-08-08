package main

import (
	"fmt"
	"os"
)

func main() {
	// Mengambil hostname (sistem operasinya)
	host, err := os.Hostname() // returnnya adalah 2 variabel

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(host)
	}
}
