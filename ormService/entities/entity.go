package entities
import (
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

var Enginea *xorm.Engine
func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
    Enginea, _ = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	err := Enginea.Ping()
	CheckErr(err)
	Enginea.SetMapper(core.SameMapper{})
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
