package ginfeng

import (
	"fmt"
	"Base"
	"Cache"
	"Config"
	"DB"
	"github.com/gin-gonic/gin"
	"reflect"
)

type WebServer struct {
	gin *gin.Engine
}

// 获取web实例
func NewWebServer() *WebServer {
	return &WebServer{
		gin: gin.Default(),
	}
}

// 初始化
func (ws *WebServer) Init() {
	config.BaseConfig.Init()
	DB.Gorm.Init(config.BaseConfig.GetMysqlConfig())     // 初始化数据库
	Cache.Redis.Init(config.BaseConfig.GetRedisConfig()) // 初始化redis
}

func (ws *WebServer) Run() {
	_ = ws.registerRouter()                              // 注册路由
	_ = ws.gin.Run(config.BaseConfig.SysConfig.HTTPAddr) // 运行gin
}

// 注册路由
func (ws *WebServer) registerRouter() error {
	if Base.Router.Mode == "General" { // 简易模式
		if generalMap := Base.Router.GetGeneralPath(); len(generalMap) > 0 {
			for _, value := range generalMap {
				for _, v := range value {
					ws.gin.HEAD(v, ws.generalHandle)    // 注册 HEAD handle
					ws.gin.GET(v, ws.generalHandle)     // 注册 GET handle
					ws.gin.POST(v, ws.generalHandle)    // 注册 POST handle
					ws.gin.OPTIONS(v, ws.generalHandle) // 注册 OPTIONS handle
				}
			}
		}
	} else if Base.Router.Mode == "Restful" { // Restful模式
		// ...
	}
	return nil
}

// 简易路由 handle
func (ws *WebServer) generalHandle(context *gin.Context) {
	path := context.Request.URL.Path
	generalCtrl := Base.Router.GetGeneral(path)
	if generalCtrl == nil {
		panic(fmt.Sprintf("get controller failed by path '%s'", path))
	}

	// 克隆结构体
	ctrlType := reflect.TypeOf(generalCtrl).Elem()
	controller, ok := reflect.New(ctrlType).Interface().(Base.ControllerInterface)
	if !ok {
		panic("controller is not ControllerInterface")
	} else {
		controller.Init(context, Base.NewResult())
	}

	// 执行 Main()
	controller.Main()
	if result := controller.Result(); result.Status != 200 {
		ws.Response(context, result)
		return
	}
}

// 返回响应
func (ws *WebServer) Response(context *gin.Context, result *Base.Result) {
	switch result.Type {
	case "String":
		context.String(result.Status, result.Msg)
	case "Json":
		context.JSON(result.Status, result.Data)
	case "Html":
		context.HTML(result.Status, result.Msg, result.Data)
	}
}
