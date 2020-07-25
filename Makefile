include .env

start:
	@bash -c "$(MAKE) -s build start-server"

build:
	@echo "  →  Building binary..."
	@go build -o ./bin/api api/main.go

start-server:
	@port=$(port) librus_username=$(librus_username) librus_password=$(librus_password) ./bin/api