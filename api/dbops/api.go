package dbops

import (
	"database/sql"
	"log"
)

// 添加用户
func AddUserCredential(loginname string, pwd string) error {
	stmtIns, err := db.Prepare("insert into users (login_name,pwd)values(?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginname, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 查找用户
func GetUserCredential(loginname string) (string, error) {
	stmtOut, err := db.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		log.Printf("GetUser Error:%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginname).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()
	return pwd, nil
}

//删除用户
func DeleteUser(loginname string, pwd string) error {
	stmtDel, err := db.Prepare("delete from users where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("DeleterUser Error:%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginname, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
