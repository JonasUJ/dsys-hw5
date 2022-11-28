all: client frontend replica discovery

bin:
	mkdir bin

client: bin
	go build -o bin/client cmd/client/*.go
	chmod +x bin/client

frontend: bin
	go build -o bin/frontend cmd/frontend/*.go
	chmod +x bin/frontend

replica: bin
	go build -o bin/replica cmd/replica/*.go
	chmod +x bin/replica

discovery: bin
	go build -o bin/discovery cmd/discovery/*.go
	chmod +x bin/discovery

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/grpc/frontend/frontend.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/grpc/replica/replica.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/grpc/discovery/discovery.proto

clean:
	rm -r bin
