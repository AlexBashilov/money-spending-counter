build:
	go build && ./booker

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

dbStartContainer:
	docker-compose up -d --build

dbDownContainer:
	docker stop $(docker ps -q)

createdb:
		docker exec -it money-spending-counter_db_1 createdb --username=root --owner=root booker

dropdb:
		docker exec -it money-spending-counter_db_1 dropdb booker

migrateup:
		migrate -path db/migration -database "postgresql://root:root@localhost:5433/booker?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://root:root@localhost:5433/booker?sslmode=disable" -verbose down

buildjaeger:
docker run -d --name jaeger \
  -p 16686:16686 \
  -p 14268:14268 \
  jaegertracing/all-in-one:1.41

stopjaeger:
	docker stop jaeger

startjaeger:
	docker start jaeger

