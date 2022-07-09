clean:
	rm pb/*.go

gen:
	protoc --plugin=protoc-gen-go=/d/software/install/Go/bin/protoc-gen-go.exe --go_out=plugins=grpc:pb --proto_path=pb pb/*.proto

server:
	go run cmd/server/*.go -port 8080

client:
	go run cmd/client/*.go -address 0.0.0.0:8080