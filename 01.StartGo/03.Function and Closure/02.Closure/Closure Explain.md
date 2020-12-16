클로저
======

클로저는 함수 바깥에 정의된 변수를 촴조하는 함수 값이다.


마치 바깥 변수를 함수 안으로 끌어들인 듯이 변수를 읽거나 쓸 수 있다.

```
func adder() func(int) int{
  sum := 0
  return func(x int) int{
    sum += x
    return sum
  }
}

func main(){
  AdderClosuer := adder() // 변수 참조
  AdderClosuer(3) // 0 + 3 = 3
  AdderClosuer(4) // 3 + 4 = 7
```
