package errs

type ErrorCode int

type BError struct {
	Code ErrorCode
	Msg  string
}

// 重写了这个方法--- 就等于它是error的类型了 ！！！！！！！！！
func (e *BError) Error() string {
	//这里可能有问题
	return e.Msg
}

func NewError(code ErrorCode, msg string) *BError {
	return &BError{
		Code: code,
		Msg:  msg,
	}
}
