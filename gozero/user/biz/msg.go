package biz

// 自定义 error
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	return e.Msg
}

const Ok = 200

var (
	DBError         = NewError(10000, "数据库错误")
	AlreadyRegister = NewError(10100, "用户已注册")
	NameOrPwdError  = NewError(10101, "账号或密码错误")
	TokenError      = NewError(10102, "token错误")
)

// 统一返回
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) *Result {
	return &Result{
		Code: Ok,
		Msg:  "success",
		Data: data,
	}
}

func Fail(err *Error) *Result {
	return &Result{
		Code: err.Code,
		Msg:  err.Msg,
	}
}
