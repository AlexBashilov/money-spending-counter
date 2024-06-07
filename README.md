[# booker
this service calculates costs and provides a report in the form of tables



Для запуска сервера выполнить следующие команды:

go build  && ./apiserver


Создаем БД
createdb booker

Миграции:

migrate --path migrations --database "postgres://postgres:post

генерация документации swagger:

swag init 

форматирование документации

swag fmt