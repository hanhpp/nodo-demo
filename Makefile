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
.PHONY: doc
doc:
	swag init -g ./api-portal/routes/stock_handler/stock.go -o ./docs
