GO_TEST_INTEGRATION_ARGS ?= -race -cover -tags="integration" -coverprofile=coverage.out -covermode=atomic

docker-build:
	docker build --build-arg project=FirebaseGo -t firebase-go -f Dockerfile .

run:
	docker-compose down
	docker-compose stop
	docker-compose build
	docker-compose up

go-test:
	go fmt ./...
	GO_ENV=test go test ./...

go-integration-test:
	docker-compose run --rm admin-test go test $(GO_TEST_INTEGRATION_ARGS) ./...

go-shell:
	docker-compose run --rm admin-test /bin/bash

