package settings

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)
var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//db, err := sql.Open("postgres", "user=micl password=miclpwd dbname=zjhmgeadata sslmode=disable")
var DB *sql.DB

func ConnectDB() {
	var err error
	DB,err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("数据库连接成功")
}