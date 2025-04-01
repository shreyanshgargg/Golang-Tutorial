package main

import "fmt"

func main() {

	fmt.Println("welcome to a class on arrays")

	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "mango"
	fruitlist[2] = "banana"
	fruitlist[3] = "grapes"

	fmt.Println("fruit list is", fruitlist)
	fmt.Println("length of fruit list is", len(fruitlist))
}
