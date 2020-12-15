package main

import "fmt"

func main() {
	var i int = 9

	var iptr = &i

	fmt.Println(*iptr) // i의 주소값 9를 저장

	var y *int = iptr // y에 i의 주소값을 가진 iptr 공유

	fmt.Println(*y) // 9출력

	*iptr = 4 // i의 주소값을 공유하는 iptr의 주소값을 4로 변경

	fmt.Println(*y) // y의 주소값도 4로 변경되어 출력
}
