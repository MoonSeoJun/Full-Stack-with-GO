package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	Sumcloser := adder()
	fmt.Println(Sumcloser(2))
	fmt.Println(Sumcloser(3))

}
