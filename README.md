[# booker
this service calculates costs and provides a report in the form of tables



Для запуска сервера выполнить следующие команды:

go build  && ./apiserver


Создаем БД
createdb booker

Миграции:

migrate create -ext sql -dir db/migrations -seq init_schema


генерация документации swagger:

swag init 

форматирование документации

swag fmt
