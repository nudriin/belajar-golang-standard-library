package main

import (
	"fmt"
	"sort"
)

// ! Membuat sort harus implement interface Len, Less, Swap pada interface sort

type User struct {
	Name string
	Age  int
}

// Kita buat type aliasnya untuk slice user
type UserSlice []User

// Implements Len interface
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] // menukar
}

func main() {
	users := []User{
		{Name: "Naruto", Age: 19},
		{Name: "Nurdin", Age: 20},
		{Name: "Elon", Age: 36},
		{Name: "Sasuke", Age: 17},
	}
	fmt.Println("BEFORE SORT :", users)

	sort.Sort(UserSlice(users)) // harus di casting ke UserSlice karena

	fmt.Println("AFTER SORT :", users)
}
