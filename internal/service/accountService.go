package service

import (
	"context"
	"fmt"
	"im-GIN/internal/dao"
	"im-GIN/internal/datastore"
	"im-GIN/internal/errors"
	"im-GIN/internal/model"
	"im-GIN/internal/utils"
	"sync"
	"time"
)

type AccountService struct{}

var (
	accountService *AccountService
	accountOnce    sync.Once
)

func GetAccountService() *AccountService {
	accountOnce.Do(func() {
		accountService = new(AccountService)
	})
	return accountService
}

func (*AccountService) Login(account *model.Account) (string, error) {
	// 判断是验证码登录还是密码登录
	var user *model.User
	var err error
	switch account.Mode {
	case 1:
		// 进入密码登录
		user, err = accountService.LoginWithPwd(account)
		if err != nil {
			return "", errors.NewServerError("帐号或密码错误", err)
		}
	//case 2:
	//	// 验证码登录 todo
	default:
		return "", errors.SimpleError("登录时发生了错误")
	}
	// 验证通过 生成jwt
	expires := time.Minute * 10
	token, err := utils.GenerateAccessToken(user.Id, expires)
	if err != nil {
		return "", errors.NewServerError("登录失败", err)
	}
	key := fmt.Sprintf("access_%d", user.Id)
	// 存入缓存
	err = datastore.Cache.Set(context.Background(), key, token, expires).Err()
	if err != nil {
		return "", errors.NewServerError("登录失败！", err)
	}
	return token, nil
}

func (*AccountService) LoginWithPwd(account *model.Account) (*model.User, error) {
	accountDao := dao.GetAccountDao()
	user, err := accountDao.LoginWithPwd(account.Phone, account.Pwd)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*AccountService) Register(user *model.User) error {
	accountDao := dao.GetAccountDao()
	err := accountDao.Register(user)
	//if err != nil {
	//	// Log
	//}
	return err
}
