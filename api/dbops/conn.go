//处理1个公用的数据库链接

package dbops

import (
	"database/sql"
	//虽然没直接使用mysql.XXX但是通过init方式使用mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// 使用全局变量，将获取到的链接存储起来，以便其他函数调用
var (
	db  *sql.DB
	err error
)

// 利用包的初始化，自动链接数据库
func init() {
	db, err = sql.Open("mysql", "mysql:123456@tcp(94.191.18.219:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
