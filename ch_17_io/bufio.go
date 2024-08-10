package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func BufioRead() {
	input := strings.NewReader("My name is Nurdin\nI am a golang and node js programmer")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine() // membaca 1 baris

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}

func BufioWrite() {
	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.Write([]byte("Nurdin hishasy is my name\n"))
	_, _ = writer.Write([]byte("I am a software engineer\n"))
	_, _ = writer.Write([]byte("I like to coding"))

	writer.Flush()
}

func main() {
	BufioRead()
	BufioWrite()
}
