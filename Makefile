build:
	go build && ./booker

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

dbStartContainer:
	docker-compose up -d --build

dbDownContainer:
	docker stop $(docker ps -q)

createdb:
		docker exec -it booker_db_1 createdb --username=root --owner=root booker

dropdb:
		docker exec -it booker_db_1 dropdb booker

migrateup:
		migrate -path db/migration -database "postgresql://root:root@localhost:5433/booker?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://root:root@localhost:5433/booker?sslmode=disable" -verbose down


