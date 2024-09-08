package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"time"
)

var DB *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("root:root@tcp(localhost:3306)/goblog?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln("数据库链接错误")
	}

	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Fatalln("数据无法连接")
		_ = DB.Close()
		panic(err)
	}
	DB = db
}
