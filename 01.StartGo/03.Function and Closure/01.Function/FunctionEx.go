package main

import "fmt"

func Add(a, b int) int { // 일반적인 리턴 함수
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func AddSubStruct(a, b int) (Add, Sub int) { // 네임드 리턴함수
	Add = a + b
	Sub = a - b
	return
}

func Execute(op func(int, int) int, a, b int) int { // 함수를 다른 함수의 인자로 사용
	return op(a, b)
}

func infinitAdder(input ...int) (sum int) { // 가변함수
	for _, v := range input {
		sum += v
	}
	return
}

func main() {
	var Adder = func(a, b int) int { // 함수를 변수에 바인드해서 사용가능
		return a + b
	}

	var result = Adder(4, 3) // 함수를 변수에 바인드

	fmt.Println(result)

	fmt.Println(Execute(Adder, 2, 4)) // Adder 호출

	fmt.Println(AddSubStruct(5, 3))

	fmt.Println(infinitAdder(53, 2, 1, 34))
}
