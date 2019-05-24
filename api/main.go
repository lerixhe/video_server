package main

import (
	"log"
	"net/http"

	//引用路由库 go get github.com/julienschmidt/httprouter
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

// 创建中间件handler，劫持原本的handler,从而实现功能
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{r}
	return m
}

// 给middlewareHandle编订ServeHttp()方法，实现了此方法才能作为Hander的参数传入listenAndServe
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session 使用validateUserSession(r)
	m.r.ServeHTTP(w, r)

}

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
	log.Println("service start @ 127.0.0.1:8000 !")
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)

}
