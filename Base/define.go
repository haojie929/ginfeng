package Base

/**************************************** 数据类型 - 结构体Result ****************************************/
// 定义常量
const (
	RespString = "String"
	RespJson   = "Json"
	RespHtml   = "Html"
)

// 响应结果
type Result struct {
	Status int         // 状态码: [200:OK] [400:Bad Request] [500:Internal Server Error] [900:逻辑异常(状态码200)]
	Type   string      // 响应类型: String、Json、Html 默认Json
	Msg    string      // 消息提示
	Data   interface{} // 响应数据
}

// 实例化 Result
func NewResult() *Result {
	return &Result{
		Status: 200,
		Type:   RespJson,
		Msg:    "",
		Data:   "",
	}
}

// 设置Code
func (r *Result) SetStatus(status int) *Result {
	r.Status = status
	return r
}

// 设置Code
func (r *Result) SetType(t string) *Result {
	r.Type = t
	return r
}

// 设置Msg
func (r *Result) SetMsg(msg string) *Result {
	r.Msg = msg
	return r
}

// 设置Data
func (r *Result) SetData(data interface{}) *Result {
	r.Data = data
	return r
}

// 抛出异常
func (r *Result) ThrowException() *Result{
	r.SetStatus(900)
	return r
}

