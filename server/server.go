package server

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"regexp"
	. "web-zjh/mainStructs"
)

var stateCode = [...]int{200, 400, 403}

//type ApiRegister struct {
//	apis []ApiInfo
//}

func handler(controller Controller, w http.ResponseWriter, r *http.Request) (err error) {
	resp, err := controller(r)
	if err != nil {
		return
	}
	//设置header
	for k, v := range resp.Header {
		w.Header().Set(k, v)
	}
	w.Header().Set("Content-Type", "application/json")
	//设置state
	if resp.State != 0 {
		for _, sc := range stateCode {
			if sc == resp.State {
				w.WriteHeader(sc)
			}
		}
	}
	//写response body
	dataByte, err := json.Marshal(resp.Data)
	if err != nil {
		log.Println("main.ControllerToHandler.json.Marshal:", err)
		return
	}
	_, err = w.Write(dataByte)
	if err != nil {
		log.Println("main.ControllerToHandler.w.Write:", err)
		return
	}
	return
}

//
//func (p *ApiRegister) Add(info ApiInfo) {
//	for _, apiInfo := range p.apis {
//		if apiInfo.Path == info.Path {
//			log.Fatal("path重复")
//		}
//	}
//
//	p.apis = append(p.apis, info)
//}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(400)
	type Data struct {
		Error string `json:"error"`
	}
	data := Data{err.Error()}
	dataByte, err := json.Marshal(data)
	if err != nil {
		log.Println("main.ServeHTTP.json.Marshal:", err, "data:", data)
	}
	_, err = w.Write(dataByte)
	if err != nil {
		log.Println("main.ServeHTTP.w.Write:", err)
	}
}

type ApiRegister []struct {
	App  string
	Urls map[string]Controller
}

//中间件在这个函数运行
func (ar *ApiRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	notFount := true
	log.Println(r.URL.Path, r.Host, r.Method)
	//可以在这里提前处理request

	// 匹配url
	for _, apps := range *ar {
		match := false
		for url, controller := range apps.Urls {
			matched, err := regexp.MatchString(path.Join(apps.App, url), r.URL.Path)
			if err != nil {
				log.Fatal("regexp.MatchString: ", err)
			}
			if matched {
				match = true
				err = handler(controller, w, r)
				if err != nil {
					handleError(err, w)
				}
				notFount = false
				break
			}
		}
		if match {
			break
		}
	}
	if notFount {
		http.NotFound(w, r)
	}
}
