package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	middle "im-GIN/api/middleware"
	"im-GIN/api/router"
	"im-GIN/config"
	_ "im-GIN/docs"
	"im-GIN/internal/datastore"
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
