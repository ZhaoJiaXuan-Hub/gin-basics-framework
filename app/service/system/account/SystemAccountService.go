package account

import (
	"admin/app/entity"
	"admin/config/database"
)

/**
新建User信息
*/
func CreateSystemAccount(user *entity.SystemAccount) (err error) {
	if err = database.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}
