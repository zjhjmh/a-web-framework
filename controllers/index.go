package controllers

import (
	"net/http"
	. "web-zjh/mainStructs"
)

func IndexHandler(r *http.Request) (resp Resp, err error) {
	resp.Data = struct {
		Code    int    `json:"code"`
		Content string `json:"content"`
	}{
		200, "hello",
	}
	resp.Header = map[string]string{
		"test": "test",
	}
	return
}

func IndexController(r *http.Request) (resp Resp, err error) {

	return
}
