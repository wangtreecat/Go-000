package dao

import (
	"Week02/model"
	"Week02/wkerror"
	"database/sql"
	"errors"

	xerrors "github.com/pkg/errors"
)

var (
	db *sql.DB
)

func GetUserNameByID(id int) (*model.User, error) {
	u, err := queryFromDB(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, wkerror.ErrUserNotFound
		}
		return nil, xerrors.Wrap(err, "GET USER NAME ERROR")
	}
	// do someting

	return u, nil
}

// 模拟两种从DB中获取数据失败
// 0 ErrNoRows
// 1 ErrConnDone
func queryFromDB(id int) (*model.User, error) {
	var u model.User
	if 0 == id {
		return &u, sql.ErrNoRows
	}
	return &u, sql.ErrConnDone
}
