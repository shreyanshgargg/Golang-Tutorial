package main

import (
	"fmt"
)

func main() {

	fmt.Println("welcome to a class on structs")

	shreyansh := User{"Shreyansh", "shreyansh@", true, 16}
	fmt.Println(shreyansh)
	fmt.Printf("Shreyansh's details are : %+v\n", shreyansh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
