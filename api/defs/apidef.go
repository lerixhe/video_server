package defs

//requests
// 用户凭证
type UserCredential struct {
	Name string `json:"user_name"`
	Pwd  string `json:"pwd"`
}

type VideoInfo struct {
	Id          string
	AuthorId    int
	Name        string
	DisplayTime string
}
type CommentInfo struct {
	Id         string
	VideoId    string
	AuthorName string //数据表中存的author_id，代码结构中使用name
	Content    string
}
type SimpleSession struct {
	UserName string
	TTL      int64 //time to live 过期时间
}
