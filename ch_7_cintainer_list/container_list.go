package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Membuat linked list
	var data *list.List = list.New()
	data.PushBack("Nurdin")  // menambah data ke belakang
	data.PushBack("Hishasy") // menambah data kebelakang
	data.PushBack("Sunny")   // menambah data kebelakang

	// MEngakses data pada list
	// balikannya adalah list element
	var element *list.Element = data.Front() // mengambil data depan ("Nurdin")

	// Mengambil value dari elementnya
	var value = element.Value
	fmt.Println(value)

	// mengambil element selanjutnya
	nextElement := element.Next()  // ("Hishasy")
	fmt.Println(nextElement.Value) //

	nextElement = nextElement.Next() // ("Sunny")
	fmt.Println(nextElement.Value)

	// Bisa di looping
	for el := data.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}

}
