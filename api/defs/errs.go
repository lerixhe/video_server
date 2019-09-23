package defs

// 自定义错误类型
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// 自定义错误回复类型
type ErroResponse struct {
	HttpSC int
	Error  Err
}

// 实例化两个错误
var (
	ErrRequestBodyParseFailed = ErroResponse{HttpSC: 400, Error: Err{Error: "Request Body Is Not Correct!", ErrorCode: "001"}}
	ErrNotAuthUser            = ErroResponse{HttpSC: 401, Error: Err{Error: "User Authentication Failed!", ErrorCode: "002"}}
	ErrDBError                = ErroResponse{500, Err{"DB ops failed", "003"}}
	ErrInternalFaults         = ErroResponse{500, Err{"Internal service faults", "004"}}
)
