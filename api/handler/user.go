package handler

import (
	"github.com/gin-gonic/gin"
	resp "im-GIN/internal/common"
	"im-GIN/internal/errors"
	"im-GIN/internal/model"
	"im-GIN/internal/service"
)

func GetUserInfo(c *gin.Context) {
	c.Set(resp.RES, resp.Success("ok"))
}

// ModifyUserInfo
// @Tags [user]
// @Summary 修改用户个人信息
// @Description 可传入任意用户个人信息进行修改
// @Accept json
// @Produce json
// @Param token header string true "Token"
// @param account body model.User true "用户信息"
// @success 200 {object} resp.R
// @success 500 {object} resp.R
// @Router /user [put]
func ModifyUserInfo(c *gin.Context) {
	var user model.User
	var err error
	if err = c.ShouldBindJSON(&user); err != nil {
		c.Error(errors.SimpleError("数据格式不对"))
		c.Abort()
		return
	}
	userService := service.GetUserService()
	if err = userService.ModifyUserInfo(&user); err != nil {
		c.Error(errors.SimpleError("修改失败"))
		c.Abort()
		return
	}
	c.Set(resp.RES, resp.Success("修改成功"))
}
