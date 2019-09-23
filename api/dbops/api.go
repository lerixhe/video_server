package dbops

import (
	"database/sql"
	"log"
	"time"
	"video_server/api/defs"

	uuid "github.com/satori/go.uuid"
)

// 添加用户，插入数据库
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

// 查找用户，from 数据库
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

//删除用户，同步到数据库
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

//添加视频,返回一个视频对象，同时写入数据库
func AddVideo(authorId int, name string) (*defs.VideoInfo, error) {
	video := defs.VideoInfo{AuthorId: authorId, Name: name}
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	video.Id = uuid.String()
	video.DisplayTime = time.Now().Format("2006-01-02 15:04:05")
	//写入数据库
	stmtav, err := db.Prepare("insert into video_info (id,author_id,name,display_time) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	stmtav.Exec(video.Id, video.AuthorId, video.Name, video.DisplayTime)
	defer stmtav.Close()
	return &video, nil
}

//删除视频
func DeleteVideo(videoId string) error {
	stmtdel, err := db.Prepare("delete from video_info where id = ?")
	if err != nil {
		return err
	}
	stmtdel.Exec(videoId)
	defer stmtdel.Close()
	return nil
}

//添加评论
func AddComment(videoId string, authorId int, content string) error {
	uuid, err := uuid.NewV4()
	commentId := uuid.String()
	//comment := defs.CommentInfo{Id: commentId, VideoId: videoId, AuthorId: authorId, Content: content}
	stmtadd, err := db.Prepare("insert into comments (id,video_id,author_id,content)values(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtadd.Exec(commentId, videoId, authorId, content)
	if err != nil {
		return err
	}
	defer stmtadd.Close()
	return nil
}

// 获取某个视频的评论列表
// from to,为分时间段获取
func GetCommentsList(videoId string, from, to int) ([]*defs.CommentInfo, error) {
	var comments []*defs.CommentInfo
	stmtgets, err := db.Prepare(`
    select comments.id,users.login_name,comments.content
    from comments join users
    on comments.author_id = users.id
    where comments.video_id = ?
    and comments.time > from_unixtime(?)
    and comments.time < from_unixtime(?)
    `)
	rows, err := stmtgets.Query(videoId, from, to)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, name, content string
		err := rows.Scan(&id, &name, &content)
		if err != nil {
			return nil, err
		}
		c := &defs.CommentInfo{
			Id:         id,
			VideoId:    videoId,
			AuthorName: name,
			Content:    content}
		comments = append(comments, c)
	}
	defer stmtgets.Close()
	return comments, nil
}
