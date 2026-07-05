postgres:
	docker run --name postgres12 \
	-p 5432:5432 \
	-e POSTGRES_USER=$(DB_USER) \
	-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
	-d postgres:12-alpine


createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bankcore

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root bankcore

migrateup:
	migrate -path ./db/migration -database "postgresql://root:secretpass@localhost:5432/bankcore?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migration -database "postgresql://root:secretpass@localhost:5432/bankcore?sslmode=disable" -verbose down

migrateup1:
	migrate -path ./db/migration -database "postgresql://root:secretpass@localhost:5432/bankcore?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path ./db/migration -database "postgresql://root:secretpass@localhost:5432/bankcore?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/Samuelmasih6/Bankcore/db/sqlc Store

.PHONY:
	createdb dropdb postges migrateup migratedown migrateup1 migratedown1 sqlc server mockdb
