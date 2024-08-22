作为 blockchain_07 的测试 SDK 使用。
# 使用方法
```
cd sdk_test
go mod tidy
go run sdk.go
```

出现 [GIN-debug] Listening and serving HTTP on :9099 说明成功在9099端口开始监听。

然后新开一个终端，在其中输入curl命令来测试。

···

curl localhost:9099/queryAllOrders
curl -v http://localhost:9099/createOrder -d '{"ID":"ORDER10","userID":6,"ticketID":"3","num":1,"totalPrice":100.00,"orderDate":"2023-04-30T12:34:56Z"}' -X POST -H "Content-Type: application/json"

···

# Q & A
Q: 为什么我出现了如下报错？
···
load MSPs from config failed: configure MSP failed: sanitizeCert failed the supplied identity is not valid: x509: certificate signed by unknown authority
···
A: 降低虚拟机的系统 go 版本到 1.18 即可