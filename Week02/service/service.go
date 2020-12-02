package service

import (
	"Week02/dao"
	"Week02/model"
)

func QueryUserInfo(id int) (*model.User, error) {
	return dao.GetUserNameByID(id)
}
