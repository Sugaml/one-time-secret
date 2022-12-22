postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple

dropdb:
	docker exec -it postgres dropdb simple

migratecreate:
	migrate create -ext sql -dir ./migrations -seq add_users

migrateup:
	migrate  -path ./migrations -database "postgresql://root:secret@localhost:5432/simple?sslmode=disable"  -verbose up 

migrateup1:
	migrate -path ./migrations -database "postgresql://root:secret@localhost:5432/simple?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./migrations -database "postgresql://root:secret@localhost:5432/simple?sslmode=disable" -verbose down

migratedown1:
	migrate -path internal/model/migration -database "postgresql://root:secret@localhost:5432/simple?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run ./cmd/server/main.go

mock:
	mockgen -package mockdb -destination internal/controller/mock/mock.go github.com/berrybytes/simplesecrets/internal/model/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
