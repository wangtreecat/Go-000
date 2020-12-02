package api

import (
	"Week02/service"
	"fmt"
)

//查询用户资料
func QueryUserInfo(userID int) {
	user, err := service.QueryUserInfo(userID)
	if err != nil {
		fmt.Printf("user not found, UserID :%d, \n%+v\n", userID, err)
	}
	fmt.Printf("user info %+v\n", user)
}
