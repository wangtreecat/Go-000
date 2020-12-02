package main

import (
	"Week02/api"
	"fmt"
)

func main() {
	api.QueryUserInfo(0)
	fmt.Println("----------\n")
	api.QueryUserInfo(1)
}
