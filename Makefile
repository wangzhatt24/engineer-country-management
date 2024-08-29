# note: call scripts from /scripts

# docker-compose.yml directory
COMPOSE_DIR=./init/

# dev dependency
dev:
	cd $(COMPOSE_DIR) && docker-compose up -d

protoc --go_out=./v1 --go_opt=paths=source_relative --go-grpc_out=./v1 --go-grpc_opt=paths=source_relative ./protos/v1/country.proto