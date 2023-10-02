run:
	go run main.go
get:
	go get -u ./...

.PHONY: docker
docker:
	docker compose -f ./docker/docker-compose.yml -p demo-db up -d
docker-down:
	docker compose -f ./docker/docker-compose.yml -p demo-db down
fmt:
	go fmt ./...
.PHONY: docs
docs:
	swag init -g ./api-portal/routes/stock_handler/stock.go -o ./docs
.PHONY: test
test:
	go test -v ./...
api:
	docker compose -f ./docker/api.yml up -d
.PHONY: migrate
migrate:
	chmod +x ./script/migrate.sh
	./script/migrate.sh