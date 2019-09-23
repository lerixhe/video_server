package session

import (
	"log"
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"

	uuid "github.com/satori/go.uuid"
)

// 定义ttl寿命为30分钟，转为毫秒
const ttlduring int64 = 30 * 60 * 1000

//这里没有使用redis之类的缓存数据库，大材小用。使用这个线程安全的map,尤其擅长并发读，但写不擅长，需要加锁
var sessionMap *sync.Map

//定义一个全局的sessionmap
func init() {
	sessionMap = &sync.Map{}
}

// 获取当前时间的毫秒数显示
func nowInMilli() int64 {
	//获取纳秒（10^-9s）转换成毫秒
	currentTime := time.Now().UnixNano() / 1000000
	return currentTime
}

//LoadSessionsFromDB 从数据库中读取session
func LoadSessionsFromDB() {
	m, err := dbops.RetrieveAllSessions()
	if err != nil {
		log.Println(err)
		return
	}
	//将得到的syncmap传给全局变量，为何不用下面的方式呢？
	// sessionMap = m
	m.Range(func(k, v interface{}) bool {
		sessionMap.Store(k, v.(*defs.SimpleSession))
		return true
	})
}

// 为用户生产新session，并返回session id
func GenerateNewSessionId(un string) string {
	//生产uuid
	id, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	sessionID := id.String()
	ttl := nowInMilli() + ttlduring
	ss := &defs.SimpleSession{un, ttl}
	//得到session和session id后将其插入syncmap中，注意还要插入数据库中
	sessionMap.Store(sessionID, ss)
	dbops.InserSession(sessionID, ttl, un)
	return sessionID
}

// 判断某个session 是否过期，过期则返回true，否则返回false和
func IsSessionExpires(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		//如果已经过期，需要清除掉此session并返回true
		if ss.(*defs.SimpleSession).TTL < nowInMilli() {
			sessionMap.Delete(sid)   //从map中删除
			dbops.DeleteSession(sid) //从数据库中删除
			return "", true
		}
		//如果还没有过期
		return ss.(*defs.SimpleSession).UserName, false
	}
	//如果没有获取到session，则认为过期，并返回空
	return "", true
}
