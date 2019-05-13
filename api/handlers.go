package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

//CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

//用户登录
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

//用户上传视频
