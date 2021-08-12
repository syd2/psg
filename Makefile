#postgres : run the postgres database
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=syd0101 -d postgres:12-alpine

#creating a new database
createdb:
	docker exec -it postgres12 createdb --username=postgres --owner=postgres psg

#droping a database
dropdb:
	docker exec -it postgres12 dropdb -U postgres psg

migrate_up:
	migrate -source file://db/migrations -database "postgresql://postgres:syd0101@localhost:5432/psg?sslmode=disable" -verbose up

migrate_down:
	migrate -source file://db/migrations -database "postgresql://postgres:syd0101@localhost:5432/psg?sslmode=disable" -verbose down
sqlc_generate:
	docker run --rm -v C:\Users\admin\Desktop\psg:/src -w /src kjconroy/sqlc generate

.PHONY: createdb dropdb migrate_up postgres sqlc_generate