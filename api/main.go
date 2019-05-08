package main

import (
	"net/http"
	//引用路由库 go get github.com/julienschmidt/httprouter
	"github.com/julienschmidt/httprouter"
)

//RegisterHandlers 注册路由
func RegisterHandlers() (router *httprouter.Router) {
	router = httprouter.New()
	router.POST("/user", CreateUser)
	return
}
func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
