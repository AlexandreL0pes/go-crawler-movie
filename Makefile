run:
	go run main.go

migrate:
	go run cmd/migrate.go

reset-run:
	aws dynamodb delete-table --table-name movies --endpoint-url http://localhost:8000 --region us-east-1 
	go run cmd/migrate.go
	go run main.go

count-movies:
	aws dynamodb scan --table-name movies --select "COUNT" --endpoint-url http://localhost:8000 --region us-east-1

create-db:
	docker run -it -v "$PWD":/src -w /src nouchka/sqlite3 movies.db
	chmod 666 movies.db