package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CreateNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}
	// Menutup file setelah di tulis dan di buka
	defer file.Close()        // menghindari memory leak
	file.WriteString(message) // menuliskan file

	return nil
}

func ReadFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)

	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var message string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}

	return message, nil
}

func AddToFile(name string, message string) error {
	// Membaca dan menambah di akhir file
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		return err
	}
	// Menutup file setelah di tulis dan di buka
	defer file.Close()        // menghindari memory leak
	file.WriteString(message) // menuliskan file

	return nil
}

func main() {
	// Bisa tambahkan extensi misal .txt .word .pdf
	// CreateNewFile("sample_log_file.txt", "This is sample log file") // filenya akan otomatis dibuatkan

	// Membaca file
	result, _ := ReadFile("sample_log_file.txt")
	fmt.Println(result)

	// Membaca dan menulis ke file
	AddToFile("sample_log_file.txt", "\nAdded a message")
}
