package handler

import (
	"github.com/gin-gonic/gin"
	resp "im-GIN/internal/common"
)

func GetUserInfo(c *gin.Context) {
	c.Set(resp.RES, resp.Success("ok"))
	//var user model.User
	//if err := c.ShouldBindQuery(&user); err != nil {
	//	fmt.Println(err)
	//	c.Error(errors.NewServerError("参数错误", err))
	//	return
	//}
	//c.Set("res", resp.Success(user))
}

//func (con *UserController) Login(c *gin.Context) {}
