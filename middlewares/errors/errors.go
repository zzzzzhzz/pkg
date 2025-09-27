package errors

type Error interface {
	Error() string
	Code() int
	SetCode(code int)
	SetMessage(message string)
}

// 自定义错误类型
type myError struct {
	code    int    // 错误码
	message string // 错误消息
}

func (e *myError) Error() string {
	return e.message
}

func (e *myError) Code() int {
	return e.code
}

func (e *myError) SetCode(code int) {
	e.code = code
}

func (e *myError) SetMessage(message string) {
	e.message = message
}

// 创建一个新的自定义错误
func NewError(code int, message string) Error {
	return &myError{code: code, message: message}
}
