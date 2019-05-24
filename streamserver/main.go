package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{r, NewConnLimiter(cc)}
	return m
}
func (mh middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !mh.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}
	mh.r.ServeHTTP(w, r)
	defer mh.l.ReleaseConn()
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploaderHandler)
	return router
}
func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 3)
	http.ListenAndServe(":9000", mh)
}
