# note: call scripts from /scripts

# docker-compose.yml directory
COMPOSE_DIR=./init/

# dev dependency
dev:
	cd $(COMPOSE_DIR) && docker-compose up -d

# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./engineer.proto