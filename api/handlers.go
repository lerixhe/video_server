package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"

	"github.com/julienschmidt/httprouter"
)

//CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//用户对路由/user发送了post请求，一定有个body
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	//将读取到的body写入到用户结构体中
	err := json.Unmarshal(res, ubody)
	if err != nil {
		//发送错误
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	err = dbops.AddUserCredential(ubody.Name, ubody.Pwd)
	if err != nil {
		sendErrorResponse(w, defs.ErrDBError)
		return
	}
	sessionID := session.GenerateNewSessionId(ubody.Name)
	//生成注册成功的reponse
	su := defs.SignedUp{true, sessionID}
	resp, err := json.Marshal(su)
	if err != nil {
		sendErrorResponse(w, defs.ErrInternalFaults)
	}
	sendNormalResponse(w, string(resp), 201)
}

//用户登录
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

//用户上传视频
