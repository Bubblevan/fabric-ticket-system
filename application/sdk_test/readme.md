��Ϊ blockchain_07 �Ĳ��� SDK ʹ�á�
# ʹ�÷���
```
cd sdk_test
go mod tidy
go run sdk.go
```

���� [GIN-debug] Listening and serving HTTP on :9099 ˵���ɹ���9099�˿ڿ�ʼ������

Ȼ���¿�һ���նˣ�����������curl���������ԡ�

������

curl localhost:9099/queryAllOrders
curl -v http://localhost:9099/createOrder -d '{"ID":"ORDER10","userID":6,"ticketID":"3","num":1,"totalPrice":100.00,"orderDate":"2023-04-30T12:34:56Z"}' -X POST -H "Content-Type: application/json"

������

# Q & A
Q: Ϊʲô�ҳ��������±���
������
load MSPs from config failed: configure MSP failed: sanitizeCert failed the supplied identity is not valid: x509: certificate signed by unknown authority
������
A: �����������ϵͳ go �汾�� 1.18 ����