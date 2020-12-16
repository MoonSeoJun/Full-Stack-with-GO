맵
====
------
+ ### 맵의 정의 
맵은 키로 해당 값을 찾을 수 있는 키-값 쌍의 조합이다.
기본적인 맵 선언은 다음과 같다.
```
var ExMap map[int]string
```
키는 int, 값은 string인 맵이다.

----------------------
+ ### 맵 초기화
```
var ExMap map[int]string

var ExMap = make(map[int]string)

var ExMap = map[int]string{key : "Value"}
```
위 3가지 방법으로 맵을 초기화할 수 있다.

---------------------
+ ### 맵 키와 값 삽입/삭제
```
var ExMap = map[int]string{1 : "First", 2 : "Second", 3 : "Third"}

ExMap[4] = "Fourth
```
위와 같이 맵에 값을 삽입할 수 있다.

```
var ExMap = map[int]string{1 : "First", 2 : "Second", 3 : "Third"}

delete(ExMap, 3)
```
위와 같은 방법으로 맵의 키와 값을 삭제할 수 있다.

-----------------------
+ ### 특정 키 찾기
```
var ExMap = map[int]string{3 : "asdf", 4 : "bdf", 5 : "eqw"}

x, ok := ExMap[4]
```
ExMap안에 4변 키가 없다면 ok는 false

존재한다면 ok는 ture, x는 해당 값이 된다.
