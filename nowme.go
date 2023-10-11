package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID   int
    Name string
}

func main() {
    // Retrieve the User profile
    u, err := retrieveUser("sally")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Display the user profile
    fmt.Printf("%+v\n", *u)
}

// 1. cool: 返回值只有两个；最好别超过三个，返回参数在标准库中你也是十分少见的
// 2. 返回参数的最佳实践应该是：error 在最后一个参数，另外再加上第二个参数。（也就是最多两个返回参数是最易读的）
// 3. (*User, error) 这里返回声明没有写为：(user *User, err error); 一个是为了便于阅读函数内部提高可读性，第二是避免出现 短式变量声明 中产生 变量阴影 shadows

func retrieveUser(name string) (*User, error) {
    r, err := getUser(name)
    // 此处 r 会在下文中继续引用 所以不适用如下的局部变量的方式写作
    if err != nil {
        return nil, err
    }

    var u User
    // err = json.Unmarshal([]byte(r), &u)
    // return &u, err  // 这里返回的是 err 读者会去判断这个err到底是什么，是判断有错误的情况还是错误为nil的情况，不利于阅读
    // 由上两行代码优化成如下行代码：
    // 将 能确定的局部变量直接封死在局部范围内，不再带入后面的code中，影响阅读
    if err := json.Unmarshal([]byte(r), &u); err != nil {
        return nil, err
    }
    // 这里直接显示 返回nil ，读者一看就知道 描述的一定是成功无错误的情况
    return &u, nil
}

func getUser(name string) (string, error) {
    result := `{"id":42, "name":"sally"}`
    return result, nil
}
