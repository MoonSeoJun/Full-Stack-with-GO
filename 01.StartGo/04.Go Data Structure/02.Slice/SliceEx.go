package main

import "fmt"

func main() {
	var mySlice = []int{1, 2, 3, 5, 6, 9}
	var subSlice = mySlice[2:4]

	fmt.Println(mySlice)
	fmt.Println(cap(mySlice))
	fmt.Println(subSlice)
	fmt.Println(cap(subSlice)) // 메모리 과다

	var BigSlice = []int{1, 58, 5, 8, 6, 68}
	var smallSlice = make([]int, 3)

	copy(smallSlice, BigSlice[2:5])

	fmt.Println(smallSlice)
	fmt.Println(cap(smallSlice)) // 메모리 과다 회피

	smallSlice = append(smallSlice, 5, 8, 2)

	fmt.Println(smallSlice)
	fmt.Println(cap(smallSlice))
}
