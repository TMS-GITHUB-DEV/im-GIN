package router

import (
	"TMS-GIN/api/handler"
	middle "TMS-GIN/api/middleware"
	resp "TMS-GIN/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// 在这里注册路由
func InitRouter(r *gin.Engine) {
	// 找不到路由，重定向
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://www.baidu.com") // fixme 重定向地址需要更改
	})
	notAuthRouter(r)
	// 需要通过登录认证的使用authRouter创建路由组
	auth := r.Group("")
	auth.Use(middle.Auth)
	authRouter(auth)
}

func authRouter(r *gin.RouterGroup) {
	userRouter := r.Group("user")
	{
		userRouter.GET("", handler.GetUserInfo)
	}

	fileRouter := r.Group("file")
	{
		fileRouter.POST("upload", handler.FileUpload)
		fileRouter.GET("download/:url", handler.FileDownload)
	}
}

func notAuthRouter(r *gin.Engine) {
	accountRouter := r.Group("account")
	{
		accountRouter.POST("login", handler.Login)
		accountRouter.POST("register", handler.Register)
	}
	testRouter := r.Group("test")
	{
		testRouter.GET("", func(c *gin.Context) {
			vv1 := c.Query("kk1")
			fmt.Println(vv1)
			c.Set(resp.RES, resp.Success("get-test"))
		})
		testRouter.POST("", func(c *gin.Context) {
			//c.Set(resp.RES, resp.Fail("post-test"))
			c.JSON(500, gin.H{})
		})
	}
	swaggerRouter := r.Group("swagger")
	{
		swaggerRouter.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
