package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Print("Welcome to maths in Golang\n")

	// Generate a random number from crypto/rand
	RandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("RandomNum:", RandomNum)
}
