REGISTRY=grbpwr
IMAGE_NAME=art-admin
VERSION=master

build: statics
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

generate:
	buf generate --path ./proto/nft/nft.proto \
	--path ./proto/auth/auth.proto

statics:
	@echo "Generating combined Swagger JSON"
	@find proto/swagger -type f -name "*.json" -exec cp {} proto/swagger \;
	@GOOS="" GOARCH="" go run proto/swagger/main.go proto/swagger > app/static/swagger/api.swagger.json
	@find proto/swagger -type f -name "*.json" -exec cp {} app/static/swagger \;


clean:
	rm -rf bin
	rm -f internal/static/static_gen.go
	rm -f static/swagger/*.json
	rm -f proto/swagger/*.json
	
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
	go install github.com/go-bindata/go-bindata/go-bindata@latest
	go install golang.org/x/text/cmd/gotext@latest