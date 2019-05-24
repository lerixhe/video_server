// Package dbops 本文件从属于dbops包，但专门定义session相关的内部基础操作
package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

//写入session到数据库
func InserSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := db.Prepare("insert into sessions (session_id,ttl,login_name)values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

//取回session,from DB
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := db.Prepare("select ttl,login_name from sessions where session_id=?")
	if err != nil {
		return nil, err
	}
	var ttl, uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	res, err := strconv.ParseInt(ttl, 10, 64)
	if err != nil {
		return nil, err
	} else {
		ss.TTL = res
		ss.UserName = uname
	}
	defer stmtOut.Close()
	return ss, nil
}

//取得全部session，from db to map
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := db.Prepare("select * from sessions")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var id, ttlstr, uname string
		err := rows.Scan(&id, &ttlstr, &uname)
		if err != nil {
			log.Println(err)
			break
		}
		ttl, err := strconv.ParseInt(ttlstr, 10, 64)
		if err != nil {
			return nil, err
		} else {
			ss := &defs.SimpleSession{uname, ttl}
			m.Store(id, ss)
			log.Printf("session id:%s,ttl:%d,name:%s\n", id, ss.TTL, ss.UserName)
		}
	}
	return m, nil
}

// 删除某个session，通过id,从数据库中
func DeleteSession(sid string) error {
	stmtOut, err := db.Prepare("delete from sessions where session_id=?")
	if err != nil {
		return err
	}
	_, err = stmtOut.Query(sid)
	if err != nil {
		return err
	}
	return nil
}
