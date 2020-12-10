**基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。**

```go 
结果输出
➜  Week03 git:(main) ✗ go run main.go
^C2020/12/10 15:37:45 signal exit
2020/12/10 15:37:45 ready to stopall
2020/12/10 15:37:45 server shutdown, addr : :8080
2020/12/10 15:37:45 server shutdown, addr : :8081
2020/12/10 15:37:46 ShutDown
```

