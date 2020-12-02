**我们在数据库操作的时候，比如dao层中遇到一个sql.ErrNoRows的时候，是否应该Wrap这个error,抛给上层？为什么，应该怎么做请写出代码？**

`DAO`层只用于获取数据，只要反应真实结果即可。`biz`层 处理dao抛来的数据，需要兼容多种数据库。
因此“DAO层中遇到一个sql.ErrNoRows的时候”，应该wrap一个预定义的sentinel error，并附带上上下文信息。

```go
//测试结果如下:
user not found, UserID :0,
user not found
user info <nil>
----------

user not found, UserID :1,
sql: connection is already closed
GET USER NAME ERROR
Week02/dao.GetUserNameByID
	/Users/gin/WorkSpace/GeekBang/Golang/HomeWork/Go-000/Week02/dao/dao.go:22
Week02/service.QueryUserInfo
	/Users/gin/WorkSpace/GeekBang/Golang/HomeWork/Go-000/Week02/service/service.go:9
Week02/api.QueryUserInfo
	/Users/gin/WorkSpace/GeekBang/Golang/HomeWork/Go-000/Week02/api/api.go:10
main.main
	/Users/gin/WorkSpace/GeekBang/Golang/HomeWork/Go-000/Week02/main.go:11
runtime.main
	/usr/local/Cellar/go/1.15/libexec/src/runtime/proc.go:204
runtime.goexit
	/usr/local/Cellar/go/1.15/libexec/src/runtime/asm_amd64.s:1374
user info <nil>
```


