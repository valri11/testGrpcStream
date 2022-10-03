
all:
	go build -o testServer ./server/...
	go build -o testClient ./client/...

comms: comms/comms.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative comms/comms.proto
