package middle

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	resp "im-GIN/internal/common"
	cusErr "im-GIN/internal/errors"
	"net/http"
)

func Resp(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Fail("服务器内部错误"))
		}
	}()
	c.Next()
	for _, err := range c.Errors {
		var serverErr cusErr.ServerError
		var validationErr validator.ValidationErrors
		switch {
		case errors.As(err.Err, &validationErr):
			c.JSON(http.StatusOK, resp.FailWithCode(400, validationErr[0].Tag()))
		case errors.As(err.Err, &serverErr):
			c.JSON(http.StatusOK, resp.FailWithCode(serverErr.Code, serverErr.Msg))
		default:
			c.JSON(http.StatusOK, resp.Fail("服务器异常"))
		}
		c.Abort()
		return
	}

	if res, ok := c.Get(resp.RES); ok {
		c.JSON(http.StatusOK, res)
	} else {
		//c.JSON(http.StatusInternalServerError, gin.H{})
	}
}
