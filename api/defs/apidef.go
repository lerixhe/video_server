package defs

//requests
// 用户凭证
type UserCredential struct {
	Name string `json:"user_name"`
	Pwd  string `json:"pwd"`
}
