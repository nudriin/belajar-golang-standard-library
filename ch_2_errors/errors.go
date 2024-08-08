package main

import (
	"errors"
	"fmt"
)

// Kebiasaan  programmer golang kalau membuat error dibuatkan dulu var errornya
var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return validationError
	}

	if id != "nurdin" {
		return notFoundError
	}

	return nil
}

func main() {
	err := GetById("sdas")

	// ! Cara cek jenis error menggunakan errors.is
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
