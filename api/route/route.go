package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"im-GIN/api/handler"
	"im-GIN/api/middleware"
	"im-GIN/internal/global/resp"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://www.baidu.com") // fixme 找不到的重定向需要更改
	})
	notAuthRoute(r)
	auth := r.Group("/")
	auth.Use(middleware.Auth)
	authRoute(auth)
}

func notAuthRoute(r *gin.Engine) {
	testRoute := r.Group("/test")
	{
		testRoute.GET("/", func(c *gin.Context) {
			c.Set(resp.RES, resp.Success("test ok"))
		})
	}

	swaggerRoute := r.Group("/swagger")
	{
		swaggerRoute.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	accountRoute := r.Group("/account")
	{
		accountRoute.POST("/login", handler.Login)
		accountRoute.POST("/register", handler.Register)
	}

	fileRoute := r.Group("/file")
	{
		fileRoute.POST("/upload", handler.FileUpload)
		fileRoute.GET("/:url", handler.FileDownload)
	}
}

func authRoute(r *gin.RouterGroup) {
	userRoute := r.Group("/user")
	{
		userRoute.GET("/:phone", handler.GetUserByPhone)
	}
}
