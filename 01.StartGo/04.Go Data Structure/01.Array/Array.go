package main

import "fmt"

func main() {
	var Array [3]int
	// var Array = [3]int{1,2,3}
	// Array := [3]int{1,2,3}

	Array[0] = 1
	Array[1] = 2
	Array[2] = 3

	ArrayLength := len(Array)

	fmt.Println(Array)
	fmt.Println(ArrayLength)
}
