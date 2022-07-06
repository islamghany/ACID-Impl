DB_DSN = postgresql://root:secret@localhost:5431/bank?sslmode=disable


## db/postgres/createdb: create new database named bank
.PHONY: db/postgres/run
db/postgres/run:
	@docker run --name bank -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


## db/postgres/createdb: create new database named bank
.PHONY: db/postgres/createdb
db/postgres/createdb:
	@echo 'Creating the bank database'
	@docker exec -it bank createdb --username=root --owner=root bank


##db/postgres/dropdb: drop the database named bank
.PHONY: db/postgres/dropdb
db/postgres/dropdb:
	@echo 'Droping the bank database'
	@docker exec -it bank dropdb bank


## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	migrate -path ./db/migrations -database ${DB_DSN} up

## db/migrations/down: apply all up database down migrations
.PHONY: db/migrations/down
db/migrations/down:
	@echo 'Running down migrations...'
	migrate -path ./db/migrations -database ${DB_DSN} down

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}


## db/sqlc/generate: build all queries
.PHONY: db/sqlc/generate
db/sqlc/generate:
	sqlc generate


## server : run the api
.PHONY: server
server:
	go run ./main.go
## test :  run all unit test in the app
.PHONY: test
test:
	CGO_ENABLED=0 go test -v -cover ./...