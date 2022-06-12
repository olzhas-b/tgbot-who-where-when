.PHONY:protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative api/api.proto  && \
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative api/api.proto

.PHONY:generate-bin
generate-bin:
	go build -o bin/client cmd/client/main.go
	go build -o bin/server cmd/server/main.go

.PHONY:clean
clean:
	rm -r bin

dockerCont := `(docker ps -qa)`
dockerVol := `(docker volume ls -q)`

down:
	docker rmi -f homework-2_client homework-2_server homework-2_migration
	docker rm -f postgres-db grpc_server grpc_client migration

up:
	docker build . -f dockerfiles/Dockerfile_server -t tg_server
	docker build . -f dockerfiles/Dockerfile_client -t tg_client
	docker run -p 8080:8080 --name grpc_server -d -it tg_server
	docker run -p 8081:8081 -it tg_client