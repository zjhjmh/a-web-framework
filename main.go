package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type apiInfo struct {
	path string
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
	//http.NotFound(w, r)
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Print("服务器启动")
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
