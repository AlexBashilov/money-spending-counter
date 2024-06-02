# booker
this service calculates costs and provides a report in the form of tables



Для запуска сервера выполнить следующие команды:

go build -v ./cmd/apiserver

./apiserver

Миграции:

migrate --path migrations --database "postgres://postgres:postgres@127.0.0.1:5432/booker?sslmode=disable" up

