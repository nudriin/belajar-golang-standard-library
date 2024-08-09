package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Now() = mendapatkan waktu sekarang
	// time.Date() = membuat waktu
	// time.Parse(layout, string) = parsing string ke waktu

	currTime := time.Now()
	fmt.Println(currTime.Local()) // get Local
	fmt.Println(currTime.Date())  // get datenya
	fmt.Println(currTime.Clock()) // get waktunya

	// (tahun, bulan, jam, menit, detik, milidetik, timezone)
	times := time.Date(2020, time.December, 19, 10, 0, 0, 0, time.Local)
	fmt.Println(times)

	// RFC3339 adalah format wwaktunya
	// membuat formatnya
	formatter := "2006-01-02 15:04:05" // ! FORMATERNYA WAJIB SEPERTI INI
	value := "2003-02-19 19:03:02"     // yyyy-MM-dd HH:mm:ss
	parsedTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parsedTime)
		fmt.Println(parsedTime.Date())
		fmt.Println(parsedTime.Year())
		fmt.Println(parsedTime.Clock())
	}

}
