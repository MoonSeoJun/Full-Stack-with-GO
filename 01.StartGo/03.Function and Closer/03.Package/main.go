package main

import (
	"fmt"
	"math/rand"

	"./adder"
)

func main() {
	result := adder.Add(3, 5)

	fmt.Println(result)

	fmt.Println("random int", rand.Intn(10))
}
