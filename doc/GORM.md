# GORM



## :books:GORM

> GORM은 Go + ORM이다.

### :bulb:Overview

Web Application은 다양한 정보를 데이터베이스로 부터 읽고 쓰게된다. 이때 데이터베이스의 내용을 Application에서 사용가능한 형태로 변환해야한다. 또한 SQL 문법에 맞도록 SQL을 만들어 주어야 데이터베이스에 데이터를 쓸 수 있다.



### :mag_right:ORM이란?

**객체 관계 매핑**(Object-relational mapping; ORM)은 데이터베이스와 객체 지향 프로그래밍 언어 간의 호환되지 않는 데이터를 변환하는 프로그래밍 기법이다. 객체 지향 언어에서 사용할 수 있는 "가상" 객체 데이터베이스를 구축하는 방법이다. 객체 관계 매핑을 가능하게 하는 상용 또는 무료 소프트웨어 패키지들이 있고, 경우에 따라서는 독자적으로 개발하기도한다.



---



### :pencil2:사용방법



ORM을 사용한 예제이다. 서버에 연결하고 DB를 자동으로 마이그레이션, 생성, 삭제, 수정, 조회등을 간단히 할 수 있다.

```
db, err := gorm.Open(mysql.Open("root:raspberry@tcp(192.168.0.100:3306)/RASPBERRY_SERVER?parseTime=true"), &gorm.Config{})
if err != nil {
    return err
}

db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(user.User{}))

db.Create(&(user.User{
    Id:       "test2",
    Password: "test2",
    Email:    "test2",
    Phone:    "test2",
}))

var u user.User

db.Take(&u)

fmt.Println(u)
```

먼저 아래와 같이 데이터베이스에 연결한다.

```
db, err := gorm.Open(mysql.Open("root:raspberry@tcp(192.168.0.100:3306)/RASPBERRY_SERVER?parseTime=true"), &gorm.Config{})
if err != nil {
    return err
}
```

뒤에 ?parseTime=true 이건 기본 테이블 생성시 날짜데이터가 있는데 이를 받아올때 기본적으로 []uint8 타입으로 가져오게 되어있다. 이걸 Time객체로 받아오기 위해 추가한 내용이다.

만약에 붙이지 않으면 다음과 같은 오류가 발생한다.

```
Scan error on column index 1, name "created_at": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
```

DB에 연결이되면 마이그레이션을 해야된다. 다음과같이 user내용을 저장할 DB를 마이그레이션 하면 자동으로 Table을 생성해 준다. (이미 존재하면 생성하지 않는다.)

```
db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(user.User{}))
```

아래와같이 User구조체를 만들고 생성하면 자동으로 데이터가 생성되어 들어간다.

```
db.Create(&(user.User{
    Id:       "test2",
    Password: "test2",
    Email:    "test2",
    Phone:    "test2",
}))
```

아래는 한개의 데이터를 조회할 수 있다.

```
db.Take(&u)

fmt.Println(u)
```

이렇게 사용하면 객체를 sql문으로 만들어 DB로 저장하고 다시 조회할 때 row데이터들을 객체로 변환하는 수고를 덜수 있고, DB Migration도 자동으로 되어 편하게 사용할 수 있다.





### 조회, 생성, 수정, 삭제



0. User 구조체

```
type User struct {
    gorm.Model
    Id       string `json:"id" gorm:"column:id"`
    Password string `json:"password" gorm:"column:password"`
    Email    string `json:"email" gorm:"column:email"`
    Phone    string `json:"phone" gorm:"column:phone"`
}
```

1. 조회하기

```
e.GET("/orm/user", func(c echo.Context) error {
    db, err := gorm.Open(mysql.Open("root:raspberry@tcp(192.168.0.100:3306)/RASPBERRY_SERVER?parseTime=true"), &gorm.Config{})
    if err != nil {
        return err
    }

    db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(user.User{}))

    var users []user.User

    db.Find(&users)

    fmt.Println(users)

    return c.JSON(http.StatusOK, users)
})
```

Find를 통해 모든 유저를 조회해서 echo.Context에 JSON을 이용하여 Json형태로 출력한다.

2. 생성하기

```
e.POST("/orm/user", func(c echo.Context) error {

    db, err := gorm.Open(mysql.Open("root:raspberry@tcp(192.168.0.100:3306)/RASPBERRY_SERVER?parseTime=true"), &gorm.Config{})
    if err != nil {
        return err
    }

    db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(user.User{}))

    u := new(user.User)
    if err = c.Bind(u); err != nil {
        return err
    }

    fmt.Println(u)

    db.Create(u)

    return c.JSON(http.StatusOK, "Created")
})
```

echo.Context에 bind를 이용하면 body에 json을 객체형태로 변환하여 받을 수 있다.

그대로 Create를 하면 DB에 저장된다.

3. 업데이트하기

```
e.PUT("/orm/user", func(c echo.Context) error {
    queryID := c.QueryParam("userid")

    db, err := gorm.Open(mysql.Open("root:raspberry@tcp(192.168.0.100:3306)/RASPBERRY_SERVER?parseTime=true"), &gorm.Config{})
    if err != nil {
        return err
    }

    db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(user.User{}))

    u := new(user.User)
    if err = c.Bind(u); err != nil {
        return err
    }

    fmt.Println("queryID: ", queryID)
    fmt.Println(u)

    db.Model(&user.User{}).Where("id=?", queryID).Updates(u)

    return c.JSON(http.StatusOK, "Updated")
})
```

QueryString에 userid를 찾아서 해당 데이터를 body에 내용으로 업데이트한다.

Model로 구조체를 지정하고 id가 queryID인 데이터를 찾아서 u(body의 내용)로 업데이트한다.

4. 삭제하기

```
e.DELETE("/orm/user", func(c echo.Context) error {
    queryID := c.QueryParam("userid")

    db, err := gorm.Open(mysql.Open("root:raspberry@tcp(192.168.0.100:3306)/RASPBERRY_SERVER?parseTime=true"), &gorm.Config{})
    if err != nil {
        return err
    }

    db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(user.User{}))

    fmt.Println("queryID: ", queryID)
    db.Unscoped().Where("id = ?", queryID).Delete(&user.User{})

    return c.JSON(http.StatusOK, "Deleted")
})
```

where로 데이터를 찾고 Delete로 객체를 적해주면 해당데이터를 삭제한다.

Unscoped를 없에면 컬럼중에 deleted_at이라는 컬럼만 시간이 업데이트 되고,

Unscoped를 넣으면 완전히 데이터를 삭제한다.
