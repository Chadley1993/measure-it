.PHONY: clean
clean:
	rm data-plane

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o data-plane server.go

.PHONY: docker
docker:
	docker build -t astro-data-plane .

