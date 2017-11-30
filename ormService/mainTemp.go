package main
import (
    "fmt"
    "time"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/go-xorm/core"
)
type UserInfo struct {
	UID        int `orm:"id,auto-inc"` //语义标签
	UserName   string
	DepartName string
	CreateAt   *time.Time
}

func main() {
    enginea, _ := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
    err := enginea.Ping()
    if err != nil {
        fmt.Println("fail!!!")  
    }
    table, _ := enginea.DBMetas()
    fmt.Println(*(table[1]))
    enginea.SetMapper(core.SameMapper{})
    userinfonew := UserInfo{200, "ZekunTest", "departHappyness", nil}
    _, err2 := enginea.Insert(userinfonew)
    if err2 != nil {
        fmt.Println("insert fail~")
    }
    UserInfoGetFromId := &UserInfo{UID:100}
    has, err3 := enginea.Get(UserInfoGetFromId)
    if err3 !=nil || !has {
        fmt.Println("search fail~")
    }
    fmt.Println(*UserInfoGetFromId);
}
