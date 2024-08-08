package main

import (
	"fmt"
	"strings"
)

func main() {
	//
	lowerString := strings.ToLower("NurdiN")
	fmt.Println(lowerString)

	upperString := strings.ToUpper("NURDInnNs")
	fmt.Println(upperString)

	// Mengahpus spasi di kiri dan kanan
	trimmedString := strings.Trim("    Nurdin      ", " ")
	fmt.Println(trimmedString)

	// Mencari string dan return valuenya adalah boolean
	searchString := strings.Contains("Nurdin adalah seorang pria dari desa kokoyashi", "pria")
	fmt.Println(searchString)

	// return nya adalah slice string
	splitString := strings.Split("Nurdin Hishasy", " ")
	for _, name := range splitString {
		fmt.Println(name)
	}

	replacedString := strings.ReplaceAll("Nurdin Hishasy Ini", "i", "o") // case sensitive
	fmt.Println(replacedString)

}
