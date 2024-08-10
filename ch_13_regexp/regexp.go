package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Membuat regexnya
	rgx := regexp.MustCompile(`n([a-z])n`) //cek apakah huruf diantara n dan n mengamndung a-z kecil

	fmt.Println(rgx.MatchString("nun")) // true karena  match
	fmt.Println(rgx.MatchString("nUn")) // false karena ada kapital
	fmt.Println(rgx.MatchString("n1n")) // false karena ada angka
}
