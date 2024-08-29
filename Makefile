.PHONY: run build

run:
	go run server/main.go dev

build:
	docker build -t ghcr.io/pirsch-analytics/tour:$(VERSION) -f build/Dockerfile .
	docker push ghcr.io/pirsch-analytics/tour:$(VERSION)
