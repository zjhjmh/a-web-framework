package mainStructs

import (
	"net/http"
)

type Resp struct {
	Data   interface{}
	State  int
	Header map[string]string
}
type Handler func(w http.ResponseWriter, r *http.Request) error
type Controller func(r *http.Request) (Resp, error)

//type ApiRegister struct {
//	app string
//	path []map[string]Controller
//}
