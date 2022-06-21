clean:
	rm pb/*.go

gen:
	protoc --plugin=protoc-gen-go=/d/workspace/golang/bin/protoc-gen-go.exe --go_out=plugins=grpc:pb --proto_path=pb pb/*.proto

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080