# генерация протобуферов
genproto:
	protoc --go_out=server/ --go_opt=paths=source_relative \
    --go-grpc_out=server/ --go-grpc_opt=paths=source_relative \
    protocols/books.proto
	protoc --go_out=client/ --go_opt=paths=source_relative \
    --go-grpc_out=client/ --go-grpc_opt=paths=source_relative \
    protocols/books.proto

# чистка протобуферов
cleanprotos:
	rm -rf ./*/protocols
	rm -rf ./*/protocols

# сборка образов
build:
	docker compose -f docker-compose.yml build

# запуск сервисов
run:
	docker compose up -d
	make logsall

# остановка сервисов и удаление 
stop:
	docker compose down

# логи
logsall:
	docker compose logs -f

# логи по клиенту
logsclient:
	docker compose logs -f client

# логи по серверу
logsserver:
	docker compose logs -f server

# логи по БД
logsdb:
	docker compose logs -f database

# дамп БД
dbbackup:
	docker exec db_mysql /usr/bin/mysqldump -u root --password=password book > db/backup.sql

# запуск тестов по серверу
testserver:
	cd server/ && go test -v

# запуск тестов по клиенту
testclient:
	cd client/ && go test -v

# запуск тестов по клиенту и серверу
testall:
	cd server/ && go test -v
	cd client/ && go test -v