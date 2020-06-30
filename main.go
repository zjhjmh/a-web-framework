package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"regexp"
	"web-zjh/service/sql"
	"web-zjh/settings"
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


type auth_group struct {
	id   int
	name string
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Print("服务器启动")
	//数据库测试
	//DB := settings.DB
	settings.ConnectDB()
	defer settings.DB.Close()
	//stmt, err := settings.DB.Prepare("INSERT INTO auth_group (id,name) VALUES ($1,$2) RETU/RNING id")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//agtest := auth_group{
	//	id:   1,
	//	name: "zjh",
	//}
	//_, err = stmt.Exec(agtest.id, agtest.name)

	//查询数据
	sql.FindUserById(1)

	aR := &apiRegister{}
	for _, url := range urls {
		aR.Add(url)
	}
	err := http.ListenAndServe(":8000", aR)
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
