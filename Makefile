.PHONY: run build deps

run:
	go run server/main.go config/config.json

build:
	docker build -t ghcr.io/pirsch-analytics/tour:$(VERSION) -f build/Dockerfile .
	docker push ghcr.io/pirsch-analytics/tour:$(VERSION)

deps:
	go get -u -t ./...
	go mod tidy
	go mod vendor
