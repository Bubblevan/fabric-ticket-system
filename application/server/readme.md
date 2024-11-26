## 1、安装Mysql数据库
```
sudo apt update
sudo apt install mysql-server
sudo systemctl enable mysql
sudo systemctl start mysql
```
然后执行一些操作（应该可以简化）
```
sudo mysql -u root
CREATE DATABASE ASG;
CREATE USER 'traveler'@'localhost' IDENTIFIED BY 'hangzhou';
GRANT ALL PRIVILEGES ON ASG.* TO 'traveler'@'localhost';
FLUSH PRIVILEGES;
USE ASG;
SHOW TABLES;
```
之后表可以由db.go自行建立

## 2、启动后端服务
```
go mod tidy
go run main.go
```