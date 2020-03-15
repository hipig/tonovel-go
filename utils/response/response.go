package response

import "fmt"

type R struct {
	Code int         `json:"code"`
	Message  string      `json:"message"`
	Data interface{} `json:"data"`
}

func (r *R) IsSuccess() bool {
	if r.Code == 0 {
		return true
	}
	return false
}

func (r *R) new() *R {
	return &R{
		Code: r.Code,
		Data: r.Data,
		Message:  errCodeMap[r.Code],
	}
}

func (r *R) String() string {
	return fmt.Sprintf("code=%d,msg=%s,data=%v", r.Code, r.Message, r.Data)
}

func (r *R) Get() *R {
	return r
}

func (r *R) WithData(data interface{}) *R {
	nr := r.new()
	nr.Data = data
	return nr
}

func (r *R) WithMsg(msg string) *R {
	nr := r.new()
	nr.Message = msg
	return nr
}

func NewR(code int, data interface{}) *R {
	return &R{
		Code: code,
		Data: data,
		Message:  errCodeMap[code],
	}
}

const (
	ERROR   = iota - 1 // -1
	SUCCESS           // 0
	FAILURE
)

const (
	NOLOGIN      = 401
	NOPERMISSION = 403
	NOFOUND      = 404
)

const (
	EMPTY_DATA = iota + 1000
	EXISTS_DATA
	NO_EXISTS_DATA
	ILLEGAL_PARAM
)

const (
	QUERY_ERROR = iota + 10000
	UPDATE_ERROR
	INSERT_ERROR
	DELETE_ERROR
)

// 常用
var (
	Error   = NewR(ERROR, nil)
	Success = NewR(SUCCESS, nil)
	Failure = NewR(FAILURE, nil)

	NoLogin      = NewR(NOLOGIN, nil)
	NoPermission = NewR(NOPERMISSION, nil)
	NofoundError = NewR(NOFOUND, nil)

	EmptyData    = NewR(EMPTY_DATA, nil)
	ExsistData   = NewR(EXISTS_DATA, nil)
	NoExsistData = NewR(NO_EXISTS_DATA, nil)
	IllegalParam = NewR(ILLEGAL_PARAM, nil)

	QueryError  = NewR(QUERY_ERROR, nil)
	UpdateError = NewR(UPDATE_ERROR, nil)
	InsertError = NewR(INSERT_ERROR, nil)
	DeleteError = NewR(DELETE_ERROR, nil)

)

var errCodeMap = map[int]string{
	ERROR:   "异常",
	SUCCESS: "成功",
	FAILURE: "失败",

	NOLOGIN:      "没有登录",
	NOPERMISSION: "没有权限",
	NOFOUND:      "404",

	EMPTY_DATA:     "数据为空",
	EXISTS_DATA:    "数据已存在",
	NO_EXISTS_DATA: "数据不存在",
	ILLEGAL_PARAM:  "参数不合法",

	QUERY_ERROR:  "查询错误",
	UPDATE_ERROR: "更新错误",
	INSERT_ERROR: "添加错误",
	DELETE_ERROR: "删除错误",
}