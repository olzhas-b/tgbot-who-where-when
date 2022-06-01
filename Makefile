test:
	sleep 5

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --experimental_allow_proto3_optional api/api.proto  && \
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional api/api.proto
