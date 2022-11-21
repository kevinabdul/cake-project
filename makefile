# network:
# 	sudo docker network create privynet
mysql-service:
	sudo docker run --network host --name mysql-service -e MYSQL_ROOT_PASSWORD=superultrasecret -e MYSQL_DATABASE=privy -p 3306:3306 mysql
.PHONY:mysql-service
cake-image:
	sudo docker build --tag cake-image ./cake-service 
.PHONY:cake-image
cake-service: cake-image
	sudo docker run --network host --name cake-service -e MYSQL_ROOT_PASSWORD=superultrasecret -e MYSQL_DATABASE=privy -e MYSQL_HOST=localhost -e MYSQL_PORT=3306 -e API_SERVER_PORT=8000 cake-image
.PHONY: cake-service
migrate-up:
	GOOSE_DRIVER=mysql GOOSE_DBSTRING="root:superultrasecret@tcp/privy?charset=utf8&parseTime=True&loc=Local" goose -dir="./migration" up
.PHONY:migrate-up
migrate-down:
	GOOSE_DRIVER=mysql GOOSE_DBSTRING="root:superultrasecret@tcp/privy?charset=utf8&parseTime=True&loc=Local" goose -dir="./migration" down
.PHONY:migrate-down
migrate-reset:
	GOOSE_DRIVER=mysql GOOSE_DBSTRING="root:superultrasecret@tcp/privy?charset=utf8&parseTime=True&loc=Local" goose -dir="./migration" reset
.PHONY:migrate-reset
migrate-status:
	GOOSE_DRIVER=mysql GOOSE_DBSTRING="root:superultrasecret@tcp/privy?charset=utf8&parseTime=True&loc=Local" goose -dir="./migration" status
.PHONY:migrate-status