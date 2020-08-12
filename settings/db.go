package settings

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var err error

const (
	host       = ""
	port       = 5432
	user       = ""
	password   = ""
	dbname     = ""
	testdbname = ""
)

var psqlInfo = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,
)
var psqlTestInfo = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, testdbname,
)

//db, err := sql.Open("postgres", "user=micl password=miclpwd dbname=zjhmgeadata sslmode=disable")
var DB *sql.DB

func ConnectDB() {
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	log.Print("数据库连接成功")
}

func ConnectTestDB() {
	DB, err = sql.Open("postgres", psqlTestInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	log.Print("数据库连接成功")
}
