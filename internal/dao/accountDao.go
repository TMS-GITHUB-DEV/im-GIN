package dao

import (
	"im-GIN/internal/datastore"
	"im-GIN/internal/model"
	"sync"
)

type AccountDao struct{}

var (
	accountDao     *AccountDao
	onceAccountDao sync.Once
)

func GetAccountDao() *AccountDao {
	onceAccountDao.Do(func() {
		accountDao = &AccountDao{}
	})
	return accountDao
}

func (*AccountDao) LoginWithPwd(phone, pwd string) (*model.User, error) {
	var user model.User
	res := datastore.DB.
		Select("id, phone, email, sex, avatar_url, created_at, updated_at").
		Where("phone = ?", phone).
		Where("pwd = ?", pwd).
		Where("deleted_at = 0").
		Take(&user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (*AccountDao) Register(user *model.User) error {
	db := datastore.DB.Create(&user)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
