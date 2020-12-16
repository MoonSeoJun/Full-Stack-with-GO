함수
========

+ ### 함수의 기본 구조
```
func Add(a int, b int){
  a + b
}
```
자료형이 같을 때는 아래와 같이 정의한다.
```
func Add(a, b int)
```
반환하는 값이 있다면 아래와 같이 정의한다.
```
func Add(a, b int) int{
  return a + b
}
```

-------------------------------------

+ ### 네임드 리턴   
네임드 리턴은 반환하는 값을 미리 함수의 헤더에 정의하는 방식이다.
```
func AddSubStruct(a,b int)(add, sub int){
  add = a + b
  sub = a - b
  return
}
```
---------------------------------------   
+ ### 바인드
GO언어에서 함수를 변수에 바인드해서 사용할 수 있다.
```
var adder = func(a, b int) int{
  return a + b
}

var result = add(5, 4)
```

---------------------------------------

+ ### 함수를 다른 함수의 인자로 사용하기
아래와 같이 Execute 함수의 인자로 adder을 사용할 수 있다.
```
func Execute(op func(int, int), a, b int) int{
  return op(a, b)
}

var adder = func(a, b int) int{
  return a + b
}

Execute(adder, 5, 2)
```

---------------------------------------

+ ### 가변 함수
가변 함수는 인자의 개수를 미리 지정하지 않은 함수이다.
```
func infinit(input ...int)(sum int){
  for _, v := range input{
    sum += v
   }
   return
}

infinit(5,8,8,4,3,8)

```
위와 같이 ...을 사용하면 인자를 계속해서 받을 수 있다.
