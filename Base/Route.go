package Base

import (
	"strings"
)

// 简易模式路由容器
type GeneralMap map[string]ControllerInterface

// 路由存储器
type RouterHandler struct {
	Mode          string				// 路由模式
	generalMap    GeneralMap			// 简易模式
	generalPath   map[string][]string	// 简易模式对应路径
}

// router对象
var Router *RouterHandler

func init() {
	Router = &RouterHandler{
		Mode:          "General",
		generalMap:    GeneralMap{},
		generalPath:   map[string][]string{},
	}
}

// 设置简易路由
func (rs *RouterHandler) General(group string, controllers GeneralMap) {
	for key, value := range controllers {
		generalKey := strings.Join([]string{"", group, key}, "/")
		rs.generalMap[generalKey] = value
	}
	if _, ok := rs.generalPath[group]; !ok {
		rs.generalPath[group] = []string{
			strings.Join([]string{"", group, ":param1"}, "/"),
			strings.Join([]string{"", group, ":param1", ":param2"}, "/"),
		}
	}
}


// 获取简易路由 relativePath数组
func (rs *RouterHandler) GetGeneralPath() map[string][]string {
	return rs.generalPath
}

// 获取简易路由对应控制器
func (rs *RouterHandler) GetGeneral(path string) ControllerInterface {
	if controller, ok := rs.generalMap[path]; ok {
		return controller
	}
	return nil
}