## 1����װMysql���ݿ�
```
sudo apt update
sudo apt install mysql-server
sudo systemctl enable mysql
sudo systemctl start mysql
```
Ȼ��ִ��һЩ������Ӧ�ÿ��Լ򻯣�
```
sudo mysql -u root
CREATE DATABASE ASG;
CREATE USER 'traveler'@'localhost' IDENTIFIED BY 'hangzhou';
GRANT ALL PRIVILEGES ON ASG.* TO 'traveler'@'localhost';
FLUSH PRIVILEGES;
USE ASG;
SHOW TABLES;
```
֮��������db.go���н���

## 2��������˷���
```
go mod tidy
go run main.go
```