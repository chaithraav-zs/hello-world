package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	log.Println("hey")

	log.Println()

	fmt.Println("My favorite number is", rand.Intn(100))
}

func hello() {
	fmt.Println("print hello")
}
func hey() {
	fmt.Println("heyyy")
}