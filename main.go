package main

import (
	"log"
	"net/http"
	"web-zjh/settings"
	"web-zjh/urls"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Print("启动")
	//数据库连接
	settings.ConnectDB()
	defer func() {
		err := settings.DB.Close()
		if err != nil {
			log.Fatal("数据库关闭失败", err)
		}
	}()
	err := http.ListenAndServe(":"+settings.Listen_port, &urls.Urls)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
