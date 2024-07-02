package model

type Account struct {
	Mode  int    `json:"mode" binding:"required"` // 暂定传入1为密码登录
	Phone string `json:"phone" binding:"required"`
	Pwd   string `json:"pwd"`
	Code  string `json:"code"`
}
