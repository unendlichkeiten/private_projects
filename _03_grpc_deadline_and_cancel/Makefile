clean:
	rm pb/*.go

gen:
	protoc --plugin=protoc-gen-go=/d/workspace/golang/bin/protoc-gen-go.exe --go_out=plugins=grpc:pb --proto_path=pb pb/*.proto

server:
	go run server/*.go -port 8081

client:
	go run client/*.go -address 0.0.0.0:8081