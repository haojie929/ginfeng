package Base

import "github.com/gin-gonic/gin"

// 控制器接口
type ControllerInterface interface {
	Init(ctx *gin.Context, result *Result) 	// 初始化
	Context() *gin.Context  	// 获取gin的Context
	BeforeMain()                // 执行之前
	Main()                  	// 动作方法
	Result() *Result                       // 控制响应(辅助用法,阻断执行直接响应异常)
}

// 控制器
type Controller struct {
	context *gin.Context
	result  *Result
}

// 控制器初始化方法
func (c *Controller) Init(context *gin.Context, result *Result) {
	c.context = context
	c.result = result
}

// 简易模式下执行方法
func (c *Controller) Context() *gin.Context {
	return c.context
}

// 简易模式 - 前置方法
func (c *Controller) BeforeMain() {

}

// 简易模式 - 执行方法
func (c *Controller) Main() {

}

// 结果方法
func (c *Controller) Result() *Result {
	return c.result
}
