package ginfeng

// 运行web服务
func RunWebServer() {
	webServer := NewWebServer()         // 获取WebServer指针
	webServer.Init()                    // WebServer初始化
	webServer.Run()                     // 运行WebServer
}