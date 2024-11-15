gen:
	protoc --go_out=. --go-grpc_out=. internal/grpc/user_service.proto

clean:
	rm internal/grpc/pb/*.go

run:
	go run cmd/user_service/main.go

new_migration:
	docker run -it --rm \
 		-v "$(shell pwd)/internal/script/migrations:/db/migrations" \
 		--network host migrate/migrate \
 		create -ext sql -dir /db/migrations -seq $(MESSAGE_NAME)


up_migration:
	docker run -it --rm \
		-v "$(shell pwd)/internal/script/migrations:/db/migrations" \
		--network host migrate/migrate \
		-path=/db/migrations \
		-database "mysql://root:mysql-db@tcp(localhost:3306)/mysql_db" up


down_migration:
	docker run -it --rm \
		-v "$(shell pwd)/internal/script/migrations:/db/migrations" \
		--network host migrate/migrate \
		-path=/db/migrations \
		-database "mysql://root:mysql-db@tcp(localhost:3306)/mysql_db" down