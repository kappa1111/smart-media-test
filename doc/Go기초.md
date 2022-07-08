# Go



## :books:Go

> Go는 2007년 구글에서 개발을 시작하여 2012년에 Go1.0을 완성했다. Go는 전통적인 컴파일, 링크 모델을 따르는 범용 프로그래밍 언어이다. 시스템 프로그래밍을 위해 개발되었고, C++, Java, Python의 장점들을 모아 만들어졌다. 또한 간결한 문법을 지향하는데 Java의 절반인 25개의 키워드만으로 프로그래밍이 가능하다.



### :bulb: 문법정리

다음과 같은 환경에서 작업한다.

```
> example
|--> main.go
```

### 

### 1. Go환경에서 프로그램 작동하기

Go는 기본적으로 package가 main인 파일에 main함수를 실행한다.

따라서 main.go에 package를 main으로 설정하고, main함수를 만들면

main의 첫줄부터 코드가 실행된다.

main.go

```
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

```
Hello
```

package : 현재 파일의 패키지를 뜻한다. (main은 디렉토리 구조와 관계없이 main 패키지를 뜻한다.)

import : 사용할 패키지(API 또는 다른 디렉토리에 내용을 사용하기위해 명시한다.)를 가져온다.

fmt.Println : fmt패키지안에 Println함수를 사용한다. Println함수는 cmd창에 문자를 출력한다.

### 

### 2. 변수 사용하기

main.go

```
package example

import "fmt"

func main() {
    var text string
    text = "Hello"
    fmt.Println(text)
}
```

```
Hello
```

```
var text string
```

var : 뒤에 변수를 선언할 때 사용한다.

text : 변수명을 의미한다.

string : 변수의 타입을 의미한다.

```
text = "Hello"
```

위에서 선언한 text 변수에 Hello라는 문자값을 넣는다.

```
fmt.Println(text)
```

text의 내용을 Println함수를 이용해 cmd창에 출력한다.

Go에서는 위에 변수 선언과 값 초기화를 동시에 할 수 있다. 따라서 아래와 같이 줄일 수 있다.

```
func main() {
    text := "Hello"
    fmt.Println(text)
}
```

### 

### 3. 조건식 사용하기 (if, switch-case)

main.go

```
package main

import "fmt"

func main() {
    a := "text1"
    b := "text2"

    if a == b {
        fmt.Println("a == b")
    } else {
        fmt.Println("a != b")
    }
}
```

```
a != b
```

a, b에 각각 'text1'과 'text2'라는 문자를 선언과 동시에 초기화 해 주었다.

if : go에서는 if조건뒤에 괄호가 없다 따라서 if와 '{'사이의 조건에 따라 뒤에 실행될 내용이 결정된다.

만약 a와 b 가 같다면 'a == b'가 출력되고 다르면 'a != b'가 출력된다.

a와 b의 값이 달라서 'a != b' 가 출력된다. 만약 조건식을 a != b로 바꾸면 a == b 가 출력된다.

이외에도

| <=  | 왼쪽 값이 작거나 같을 때 |
| --- | -------------- |
| >=  | 왼쪽 값이 크거나 같을 때 |
| <   | 왼쪽 값이 작을 때     |
| >   | 왼쪽 값이 클 때      |

위와같이 사용할 수 있다.

main.go

```
package main

import "fmt"

func main() {
    a := "text1"

    switch a {
    case "text1":
        fmt.Println("1번")
    case "text2":
        fmt.Println("2번")
    default:
        fmt.Println("default")
    }
}
```

```
1번
```

switch : switch에도 괄호가 없다. switch와 '{'사이에 체크할 변수를 적고, 중괄호안의 case들과 비교해 뒤에 실행할 내용을 결정한다.

case : 변수a와 비교할 값들을 적는다. a == "text1"일 때 뒤에 fmt.Println("1번")가 실행 된다.

default : 위 케이스 조건과 모두 맞지 않을 때 default가 실행된다.

### 

### 4. 함수 사용하기

main.go

```
package main

import "fmt"

func main() {
    value1 := 10
    value2 := 20

    value3 := Add(value1, value2)

    fmt.Println(value3)
}

func Add(v1 int, v2 int) int {
    return v1 + v2
}
```

```
30
```

value1, value2 : 두 변수에 10과 20을 각각 선언과 동시에 초기화 해준다.

Add함수에 value1과 value2를 넣어 주고 value3을 선언과 동시에 초기화 해준다.

즉, Add(value1, value2)의 리턴값을 value3에 바로 값을 넣어주게 된다.

func : 함수를 선언할 때 사용한다.

Add : 함수 이름을 의미한다. 괄호안의 내용은 입력될 값을 의미하는데 

 첫번째 값은 int형이고 함수안에서 v1이란 이름으로 사용한다는 의미이다.

 두번째 값은 int형이고 함수안에서 v2이란 이름으로 사용한다는 의미이다.

) int { : 괄호와 중괄호사이의 타입은 되돌려줄 값의 타입을 의미한다.

return : 뒤의 내용을 함수를 사용한곳으로 돌려준다.

main.go

```
package main

import "fmt"

func main() {
    value1 := 10
    value2 := 20

    var value3 int

    Add(&value1, &value2, &value3)

    fmt.Println(value3)
}

func Add(v1 *int, v2 *int, v3 *int) {
    *v3 = *v1 + *v2
}
```

```
30
```

위 내용은 func을 call-by-reference방식으로 사용했다.

value1, value2를 선언과 동시에 값을 초기화 해준다.

결과를 받들 value3변수를 선언해 준다.

Add함수에 value1, value2, value3의 주소값을 넣어준다.

그리고 value3에 값이 들어와 출력된다.

func : v1, v2, v3를 int의 주소값 형태로 받아온다. 또한 리턴타입이 없다.

v1과 v2에 *로 접근하여 값을 더하고 v3에 *을 사용하여 해당 실제 변수에 값을 할당해 준다.

따라서 위에 value3에 값이 바뀌게 된다.

### 

### 5. 예외처리 하기

```
package main

import (
    "errors"
    "fmt"
)

func main() {
    value1 := 10
    value2 := 20

    value3, err := Add(value1, value2)
    if err != nil {
        fmt.Println("err: ", err.Error())
        return
    }

    fmt.Println(value3)
}

func Add(v1 int, v2 int) (int, error) {
    result := v1 + v2
    if result >= 30 {
        return 0, errors.New("결과 값이 30보다 큽니다.")
    } else {
        return result, nil
    }
}
```

```
err:  결과 값이 30보다 큽니다.
```

golang은 Exception을 error를 리턴하는 방식으로 처리됩니다.

value1과 vlaue2를 각각 선언과 초기화를 해줍니다.

Add함수를 호출하는데 이때 리턴되는 값이 두개로 하나는 int값이고 하나는 error가 된다.

Add함수를 보면 v1과 v2의 값을 받고 리턴타입에 (int, error)로 두개를 리턴하도록 되어있다. 이때 리턴타입은 괄호로 묶어주게 된다.

함수를 보면

```
result := v1 + v2
    if result >= 30 {
        return 0, errors.New("결과 값이 30보다 큽니다.")
    } else {
        return result, nil
    }
```

v1과 v2를 더하여 result변수에 선언해주고 조건식을 통해 30보다 클때 에러를 발생시킨다.

errors패키지에 New를 통해 error를 인스턴수화 해준다. 이는 객체를 실제 값으로 만들어 주는 작업을 뜻한다.

error에 메시지를 넣어 초기화 해주면 error를 처리할 때 Error()함수를 사용하게 되는데, 여기 값으로 메시지가 나온다.

만약 30보다 작으면 result를 정상적으로 리턴해주고, error는 nil을 리턴하여 함수를 상용한데서 error발생 이 없다는 것을 알린다.

다시 메인을 보면

```
if err != nil {
  fmt.Println("err: ", err.Error())
  return
}
```

err가 nil인지 확인하는데 err에 nil이 있으면 에러가 발생하지 않았고, nil이 아니면 에러가 발생한 의미가 된다.

따라서 err != nil일 때 err에 Error()함수를 사용하여 메시지를 출력 하게된다.





### :building_construction:Struct

다음과 같은 환경에서 작업한다.

```
> example
|--> main.go
```

### 1. Struct 사용하기

main.go

```
package main

import (
    "fmt"
)

type User struct {
    UserID   string
    Password string
}

func main() {
    var user User
    user = User{
        UserID:   "아이디",
        Password: "패스워드",
    }

    fmt.Println(user)
}
```

```
{아이디 패스워드}
```

type User struct : 구조체를 선언하는 방식이다 User는 구조체 이름을 뜻한다.

중괄호안에는 구조체의 내용이 들어가는데 변수의 이름과 타입을 순서대로 적는다.

```
var user User
```

var을 통해 변수선언을 해주는데 타입을 User로 한다.

```
user = User{
    UserID:   "아이디",
    Password: "패스워드",
}
```

user변수에 값을 할당할 때는 다음과같이 타입(구조체 이름)을 쓰고,

중괄호안에 각 변수들에 들어갈 값을 넣어 준다.

Go에서는 다음과 같이 선언과 동시에 초기화를 할 수있다.

```
user := User{
        UserID:   "아이디",
        Password: "패스워드",
    }
```

### 2. 구조체에 함수 추가하기

main.go

```
package main

import (
    "fmt"
)

type User struct {
    UserID   string
    Password string
}

func (user User) getUserID() string {
    return user.UserID
}

func main() {
    user := User{
        UserID:   "아이디",
        Password: "패스워드",
    }

    fmt.Println(user.getUserID())
}
```

```
아이디
```

구조체에 변수뿐만 아니라 객체처럼 메소드(함수)를 추가해 줄 수 있다.

```
func (user User) getUserID() string {
    return user.UserID
}
```

func과 getUserID(함수이름) 사이에 '(user User)' 가 있다.

User만 써도 되는데 이는 함수가 들어가 구조체를 의미한다. 따라서 User구조체를 만들어서 이 구조체를 통해서 함수를 사용할 수있다.

User 앞에 user(변수명)는 이 함수를 사용하려는 구조체를 변수로 받아온것이다.

따라서 이 함수를 사용할 구조체의 내용을 사용할 수 있게된다. (user.UserID)

메인으로 와서

```
user := User{
        UserID:   "아이디",
        Password: "패스워드",
    }

fmt.Println(user.getUserID())
```

여기서 user.getUserID를 했는데 user의 UserID와 Password가 함수안의 user변수로 받을 수 있게되는것이다.

따라서 return된 UserID값이 fmt.Println에 출력이 된다.


