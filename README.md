# go-struct-comparator

[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/louis296/go-struct-comparator)  

go语言结构体比较器
## 使用方法  
### 安装获取
```sh
go get -u  github.com/louis296/go-struct-comparator
```

### 最佳实践
```go
package main

import (
	"fmt"
	gsc "github.com/louis296/go-struct-comparator"
)

type User struct {
	Name string `compare_key:"姓名"`
	Age  int
}

type Manager struct {
	User  User   `compare_key:"用户"`
	Title string `compare_key:"标题"`
}

func main() {
	a:= Manager{User: User{Name: "wxl", Age: 1}, Title: "1"}
	b:= Manager{User: User{Name: "zzw", Age: 1}, Title: "2"}
	fmt.Println(gsc.Compare(a,b)) //map[标题:[2 1] 用户-姓名:[wxl zzw]]
}
```
### 注意事项
* 结构体之间的比较可以跨越类型，比如：
```go
type Person struct {
	Age    int     `compare_key:"年龄"`
	Name   string  `compare_key:"姓名"`
}

type PersonB struct {
	Age    int     `compare_key:"年龄"`
	Name   string  `compare_key:"姓名"`
}
```

* 嵌套的结构体之间的比较是递归的，其比较的key会以 "-" 连接
* 数组间的比较是按下标比较的，下标会以 "-" 连接在结果的key中
