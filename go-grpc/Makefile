create:
	protoc --go_opt=paths=source_relative --go_out=gen/ proto/*.proto
	protoc --go-grpc_opt=paths=source_relative --go-grpc_out=gen/ proto/*.proto

clean:
	rm gen/proto/*.go
	