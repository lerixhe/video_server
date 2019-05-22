package session

import (
	"log"
	"sync"
	"video_server/api/dbops"

	uuid "github.com/satori/go.uuid"
)

//这里没有使用redis之类的缓存数据库，大材小用
var sessionMap *sync.Map //线程安全的map,尤其擅长并发读，但写不擅长，需要加锁
func init() {
	sessionMap = &sync.Map{}
}

//LoadSessionsFromDB 从数据库中读取session
func LoadSessionsFromDB() {
	m, err := dbops.RetrieveAllSessions()
	if err != nil {
		log.Println(err)
		return
	}
	// m.Range(func(k, value interface{}) bool {
	// 	return true
	// })
}
func GenerateNewSessionId(un string) string {
	id, err := uuid.NewV4()
	if err != nil {
		return ""
	}
}
func IsSessionExpires(sid string) (string, bool) {

}
