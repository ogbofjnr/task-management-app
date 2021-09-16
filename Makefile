

all: deps lint test build


deps:
	go mod download

build:
	go build main.go

migrate:
	./migrate -path migrations -database "postgresql://root:root@127.0.0.1:5432/pm?sslmode=disable" -verbose up

format-go-code:
	gofmt -s -w .

lint:
	@command -v golangci-lint >/dev/null 2>&1 || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b /usr/local/bin v1.41.1
	golangci-lint run  --timeout=15m

test:
	go test ./... -v -timeout 30s -count=1 -p=1 -race -cover


reset-db:
	sudo docker-compose rm -fs postgres
	sudo docker-compose up -d
