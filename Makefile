test:
	sleep 5

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --experimental_allow_proto3_optional api/api.proto  && \
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional api/api.proto


run: generate-bin
	bin/server
	sleep 5
	bin/client

generate-bin:
	go build -o bin/client cmd/client/main.go
	go build -o bin/server cmd/server/main.go

.PHONY:clean
clean:
	rm -r bin