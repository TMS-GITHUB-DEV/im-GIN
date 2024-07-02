package main

import (
	middle "TMS-GIN/api/middleware"
	"TMS-GIN/api/router"
	"TMS-GIN/config"
	_ "TMS-GIN/docs"
	"TMS-GIN/internal/datastore"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 注册中间件
	middle.InitMiddleware(r)
	// 注册路由
	router.InitRouter(r)
	r.Static("static", "./static")
	// 初始化缓存
	// 初始化定时任务
	// 启动后台任务
	// 设置全局变量
	r.Run(fmt.Sprintf("%s:%s", config.Cfg.App.Host, config.Cfg.App.Port))
}

func init() {
	// 读取配置文件
	config.InitConfig()
	// 初始化日志	还不知道怎么按日期和类型存文件
	// 连接数据库
	datastore.InitDB(config.Cfg.Db)
	datastore.InitRedis(config.Cfg.Redis)
}
