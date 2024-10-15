gen:
	protoc --go_out=. --go-grpc_out=. internal/grpc/user_service.proto

clean:
	rm internal/grpc/pb/*.go