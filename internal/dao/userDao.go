package dao

import (
	"im-GIN/internal/datastore"
	"im-GIN/internal/model"
	"sync"
)

type UserDao struct{}

var (
	userDao     *UserDao
	onceUserDao sync.Once
)

func GetUserDaoInstance() *UserDao {
	onceUserDao.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (*UserDao) ModifyUserInfo(user *model.User) error {
	db := datastore.DB.Updates(&user)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
