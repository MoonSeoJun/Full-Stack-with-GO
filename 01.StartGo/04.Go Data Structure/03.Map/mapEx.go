package main

import "fmt"

func main() {
	var MapEx map[int]string

	MapEx = make(map[int]string)

	MapEx = map[int]string{1: "first", 2: "second", 3: "third"}

	fmt.Println(MapEx)

	MapEx[4] = "fourth"

	fmt.Println(MapEx)

	x, ok := MapEx[5] // 요소가 있으면 x == true, 없으면 ok = false

	fmt.Println(x)
	fmt.Println(ok)

	delete(MapEx, 3) // map 삭제

	fmt.Println(MapEx)
}
