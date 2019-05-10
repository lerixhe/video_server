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
	router.POST("/user/:user_name", Login)
	//router.POST("/user/:user_name/videos/:video_name", AddVideo)
	return
}

//主函数
func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
