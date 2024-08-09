package main

import (
	"fmt"
	"time"
)

func main() {
	oneMinutes := time.Second * 60    // 60 detik
	tenMilis := time.Millisecond * 10 // 10 milidetik
	duration := oneMinutes - tenMilis

	fmt.Println(duration)
}
