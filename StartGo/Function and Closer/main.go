package main

import "fmt"

var adder = func(a, b int) int {
	return a + b
}

var minuser = func(a, b int) int {
	return a - b
}

func execute(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func main() {
	fmt.Println(execute(adder, 3, 2))
	fmt.Println(execute(minuser, 5, 3))
}
