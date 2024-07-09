package service

import (
	"im-GIN/internal/dao"
	"im-GIN/internal/model"
	"sync"
)

type UserService struct {
}

var (
	userService *UserService
	userOnce    sync.Once
)

func GetUserService() *UserService {
	userOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (UserService *UserService) ModifyUserInfo(user *model.User) error {
	userDao := dao.GetUserDaoInstance()
	err := userDao.ModifyUserInfo(user)
	return err
}
