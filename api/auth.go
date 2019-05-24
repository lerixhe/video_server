// 此文件属于main包，主要集中了实现用户验证相关的函数
package main

import (
	"net/http"
	"video_server/api/session"
)

// 定义两个字符串：自定义header
var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

//验证用户session是否可用，同时实现
func validateUserSession(r *http.Request) bool {
	//从请求头中获取session
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpires(sid)
	if ok {
		return false
	}
	//如果session是可用的，则将用户名写如请求头并返回true
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

// 验证用户名
func validateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		return false
	}
	return true
}
