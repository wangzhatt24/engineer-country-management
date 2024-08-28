# note: call scripts from /scripts

# docker-compose.yml directory
COMPOSE_DIR=./init/

# dev
dev:
	cd $(COMPOSE_DIR) && docker-compose up -d
