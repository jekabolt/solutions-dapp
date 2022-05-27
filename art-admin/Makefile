REGISTRY=grbpwr
IMAGE_NAME=art-admin
VERSION=master

build:
	go build -o ./bin/$(IMAGE_NAME) ./cmd/

run: build
	./bin/$(IMAGE_NAME)

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	rm coverage.out

local: build
	source .env && ./bin/$(IMAGE_NAME)

image:
	docker build -t $(REGISTRY)/${IMAGE_NAME}:$(VERSION) .

image-run:
	docker run --publish 8081:8081 --env-file .env \
	--mount src=/root/bunt,target=/root/bunt,type=bind \$(REGISTRY)/${IMAGE_NAME}:$(VERSION)
