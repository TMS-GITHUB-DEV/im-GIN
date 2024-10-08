package handler

import (
	"github.com/gin-gonic/gin"
	"im-GIN/internal/global/errs"
	"im-GIN/internal/global/resp"
	"im-GIN/internal/model"
	"im-GIN/internal/service"
)

// Login
// @Tags [account]
// @Summary 登录
// @Description 登录并获取token
// @Accept json
// @Produce json
// @param account body model.Account true "登录信息"
// @success 200 {object} resp.R "登录成功"
// @success 500 {object} resp.R "登录失败"
// @Router /account/login [post]
func Login(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.Error(errs.NewServerErr("登录异常", err))
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

// LoginWithCode
// @Tags [account]
// @Summary 登录
// @Description 登录并获取token
// @Accept json
// @Produce json
// @param account body model.Account true "登录信息"
// @success 200 {object} resp.R "登录成功"
// @success 500 {object} resp.R "登录失败"
// @Router /account/login_with_code [post]
func LoginWithCode(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.Error(errs.NewServerErr("登录异常", err))
		c.Abort()
		return
	}

	accountService := service.GetAccountService()
	// fixme替换验证码登录方法
	token, err := accountService.Login(&account)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.Set(resp.RES, resp.Success(token))
}

func CMSCode(c *gin.Context) {
	c.Set(resp.RES, resp.Success(nil))
}

// Register
// @Tags [account]
// @Summary 注册
// @Description 注册账号
// @Accept json
// @Produce json
// @param user body model.User true "注册信息"
// @success 200 {object} resp.R "注册成功"
// @success 500 {object} resp.R "注册失败"
// @Router /account/register [post]
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(errs.NewServerErr("注册异常", err))
		c.Abort()
		return
	}
	accountService := service.GetAccountService()
	if err := accountService.Register(&user); err != nil {
		c.Error(errs.NewServerErr("注册失败", err))
		c.Abort()
		return
	}

	// fixme 注册后返回的信息需要更改
	c.Set(resp.RES, resp.Success(user))
}
