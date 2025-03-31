package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("value of pointor is", ptr)
	fmt.Println("value of pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new value of my number is", myNumber)
}
