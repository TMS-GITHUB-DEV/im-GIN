package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"im-GIN/internal/global/errs"
	"im-GIN/internal/global/logger"
	"im-GIN/internal/global/resp"
)

func Resp(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatus(500)
		}
	}()
	c.Next()

	for _, err := range c.Errors {
		var serverErr errs.ServerErr
		switch {
		case errors.As(err, &serverErr):
			logger.Logger.Error(fmt.Sprintf("%s：%v", serverErr.Msg, serverErr.Err.Error()))
			c.JSON(200, resp.FailWithCode(serverErr.Code, serverErr.Msg))
		default:
			logger.Logger.Error(err.Error())
			c.JSON(200, resp.Fail("服务器异常"))
		}
	}

	if res, ok := c.Get(resp.RES); ok {
		c.JSON(200, res)
	}
}
