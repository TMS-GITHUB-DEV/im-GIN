package handler

import (
	resp "TMS-GIN/internal/common"
	"TMS-GIN/internal/errors"
	"TMS-GIN/internal/model"
	"TMS-GIN/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Login
// @Tags [account]
// @Summary 登录
// @Description 登录并获取token，传入mode代表登录方式：1 密码登录
// @Accept json
// @Produce json
// @param account body model.Account true "登录信息"
// @success 200 {object} resp.R
// @success 500 {object} resp.R
// @Router /account/login [post]
func Login(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		fmt.Printf("%v", err)
		c.Error(errors.NewServerError("登录异常", err))
		c.Abort()
		return
	}

	if (account.Pwd == "") && (account.Code == "") {
		c.Error(errors.SimpleError("请使用验证码或者密码登录"))
		c.Abort()
		return
	}

	accountService := service.GetAccountService()
	token, err := accountService.Login(&account)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.Set(resp.RES, resp.Success(token))
}

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.Error(errors.NewServerError("参数错误", err))
		return
	}
	// 日志
	accountService := service.GetAccountService()
	err := accountService.Register(&user)
	if err != nil {
		c.Error(errors.NewServerError("注册失败", err))
		c.Abort()
		return
	}
	c.Set(resp.RES, resp.Success(user))
}
