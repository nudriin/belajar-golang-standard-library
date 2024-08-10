package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func Base64Encode() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Nurdin Hishasy")) // harus byte slice
	fmt.Println(encoded)

	// ! Mendecode stringnya
	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded)) // di parse ke string
	}
}

func CsvRead() {
	csvString := "Nurdin,Hisahsy,Sunny\n" + "Summer,Sasuke,Naruto\n" + "Gaara,Pain,Itachi"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		records, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(records)
		for _, data := range records {
			fmt.Println(data)
		}
	}
}

func CsvWrite() {
	writer := csv.NewWriter(os.Stdout) // akan di write ke terminal

	// ini sebenarnya akan return error
	_ = writer.Write([]string{"Nurdin", "Hishasy", "Sunny"})
	_ = writer.Write([]string{"React", "Golang", "Typescript"})
	_ = writer.Write([]string{"NestJS", "ExpressJS", "PHP"})

	writer.Flush() // me writenya
}
func main() {
	Base64Encode()
	fmt.Println()
	CsvRead()
	fmt.Println()
	CsvWrite()
}
