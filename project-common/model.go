package common

type BusinessCode int
type Result struct {
	Code BusinessCode
	Msg  string
	Data any `json:"data"`
}

// 成功  设计到取名字的地方 都要想到是否需要大写，一般都是需要大写的
func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}

func (r *Result) Fail(code BusinessCode, msg string) *Result {
	r.Code = code
	r.Msg = msg
	return r
}
