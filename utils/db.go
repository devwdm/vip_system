package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

/*
定义常量 隐
数据库配置
*/
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "vip"
	//设置数据库最大连接数
	connMax = 100
	//设置上数据库最大闲置连接数
	MaxIdleConns = 10
)

//Db数据库连接池
var (
	Db  *sql.DB
	err error
)

func init() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//导入mysql 数据库驱动： _ "github.com/go-sql-driver/mysql"
	Db, err = sql.Open("mysql", path)
	Db.SetConnMaxLifetime(connMax)
	Db.SetMaxIdleConns(MaxIdleConns)
	//验证连接
	if err != nil {
		panic(err.Error())
	}
}
