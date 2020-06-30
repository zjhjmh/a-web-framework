package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"regexp"
)

type apiInfo struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
}

type apiRegister struct {
	apis []apiInfo
}

func (p *apiRegister) Add(info apiInfo) {
	for _, apiInfo := range p.apis {
		if apiInfo.path == info.path {
			log.Fatal("path重复")
		}
	}
	p.apis = append(p.apis, info)
}

func (p *apiRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	notFount := true
	for _, apiInfo := range p.apis {
		matched, err := regexp.MatchString(apiInfo.path, r.URL.Path)
		if err != nil {
			log.Fatal("regexp.MatchString: ", err)
		}
		if matched {
			apiInfo.handler(w, r)
			notFount = false
		}
	}
	if notFount {
		http.NotFound(w, r)
	}
}

const (
	host     = "47.114.110.168"
	port     = 5432
	user     = "micl"
	password = "miclpwd"
	dbname   = "zjhmgeadata"
)

type auth_group struct {
	id   int
	name string
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Print("服务器启动")
	//数据库测试
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//db, err := sql.Open("postgres", "user=micl password=miclpwd dbname=zjhmgeadata sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("数据库连接成功")

	stmt, err := db.Prepare("INSERT INTO auth_group (id,name) VALUES ($1,$2) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}
	agtest := auth_group{
		id:   1,
		name: "zjh",
	}
	res, err := stmt.Exec(agtest.id, agtest.name)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
	aR := &apiRegister{}
	for _, url := range urls {
		aR.Add(url)
	}
	err = http.ListenAndServe(":8000", aR)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello zjh!")
	if err != nil {
		log.Fatal("indexHandler: ", err)
	}
}
