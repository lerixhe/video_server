package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
