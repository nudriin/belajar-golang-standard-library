package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
	Age  int
}

// Membuat struct tag
type StructTag struct {
	Name string `required:"true" min:"3"`
}

// ! Membuat validation dengan struct tag
type UserRequest struct {
	Email    string `required:"true" min:"3"`
	Password string `required:"true" min:"3"`
}

func IsValid(request any) (result bool) {
	result = true
	t := reflect.TypeOf(request)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("required") == "true" {
			// Di konversi ke interface
			result = reflect.ValueOf(request).Field(i).Interface() != ""
			if result == false {
				return result
			}
			// return isNotBlank
		}
	}
	return result
}

func main() {
	sample := Sample{Name: "Nurdin", Age: 20}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.Field(1).Name)

	sampleTypeName := reflect.TypeOf(sample).Name()
	fmt.Println("Type name =", sampleTypeName)
	for i := 0; i < sampleType.NumField(); i++ {
		fmt.Println("Field", sampleType.Field(i).Name, "with type", sampleType.Field(i).Type)
	}

	// mengambil tag
	user := StructTag{
		Name: "Nu",
	}
	userType := reflect.TypeOf(user)
	userField := userType.Field(0)
	tagMin := userField.Tag.Get("min")
	fmt.Println(tagMin) // 3

	// ! MEMBUAT VALIDATION
	request := UserRequest{
		Email:    "His",
		Password: "Nurdin",
	}

	fmt.Println(IsValid(request))
}
