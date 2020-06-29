package main

import (
	"fmt"
	"log"
	"net/http"
)

type apiInfo struct {
	path string
	handler func(w http.ResponseWriter, r *http.Request)
}

type apiRegister struct {
	apis []apiInfo
}

func (p *apiRegister) Add(info apiInfo) {
	//还需要判断是否重复
	p.apis = append(p.apis, info)
}

func (p *apiRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, apiInfo := range p.apis {
		if r.URL.Path == apiInfo.path {
			apiInfo.handler(w, r)
		}
	}
	//http.NotFound(w, r)
	//return
}

func main() {
	aI := apiInfo{
		path:    "/zjh",
		handler: indexHandler,
	}
	aR := &apiRegister{}
	aR.Add(aI)
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
