run:
	go run main.go
get:
	go get -u ./...

.PHONY: docker
docker:
	docker compose -f ./docker-compose.yml up -d
test:
	go test -v ./...
fmt:
	go fmt ./...
.PHONY: docs
docs:
	swag init -g ./api-portal/routes/stock_handler/stock.go -o ./docs
.PHONY: test
test:
	go test -v ./...
