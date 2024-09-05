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
      -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
      -p 5775:5775/udp \
      -p 6831:6831/udp \
      -p 6832:6832/udp \
      -p 5778:5778 \
      -p 16686:16686 \
      -p 14268:14268 \
      -p 14250:14250 \
      -p 9411:9411 \
      jaegertracing/all-in-one:1.24

stopjaeger:
	docker stop jaeger

startjaeger:
	docker start jaeger

