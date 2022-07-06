REGISTRY=jekabolt
VERSION=master

build: statics
	go build -a $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o ./bin/$(IMAGE_NAME) ./cmd/	

run: build
	./bin/$(IMAGE_NAME)

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	rm coverage.out

local: build
	source .env && ./bin/$(IMAGE_NAME)

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
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install golang.org/x/text/cmd/gotext@latest


# configuration for image names
BUILD_TYPE					:= dev # for the autobuilds will be "release"

VERSION_CURRENT_BRANCH 		:= $(shell git rev-parse --abbrev-ref HEAD)
VERSION_CURRENT_DATE 		:= $(shell date +'%Y.%m.%d')
VERSION_GIT_COMMIT 			:= $(shell git rev-parse --short HEAD)
VERSION_BUILD_NUMBER 		:= $(shell git rev-list --count HEAD) 
VERSION_TAG_LONG 			:= v_._._-$(VERSION_BUILD_NUMBER:-=)-$(VERSION_GIT_COMMIT)

GIT_TAG						:= $(shell git describe 2>/dev/null)
ifeq ($(GIT_TAG),)
	GIT_TAG					:= $(VERSION_TAG_LONG)
endif

IMAGE_VERSION 				:= $(shell printf '%s-%s-%s' $(GIT_TAG) $(VERSION_CURRENT_DATE) $(VERSION_CURRENT_BRANCH) )

# configuration for binary and image
IMAGE_NAME					:= art-admin
IMAGE_DOCKERFILE			:= $(DOCKERFILE_PATH)/Dockerfile
DOCKER_IMAGE				:= $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_VERSION)
# DOCKER_IMAGE				:= $(IMAGE_NAME):$(IMAGE_VERSION)

image-build: ## build the docker image
	docker build \
		--progress=plain \
		-t $(DOCKER_IMAGE) .


image-run:
	docker stop solutions-dapp \
    docker rm solutions-dapp \
    docker run \
			  --name=solutions-dapp -d \
              --restart=unless-stopped  \
              --publish 8001:8001 \
              --env-file .env \
              --mount src=/root/bunt,target=/root/bunt,type=bind \
              $(DOCKER_IMAGE)

test: generate
	# we only have non-generated code in ./internal, so we only count coverage for it
	go test -cover -coverprofile coverage.out -coverpkg ./app/...,./internal/... ./...
	# IMPORTANT: required coverage can only be increased
	go tool cover -func coverage.out | \
		awk 'END { print "Coverage: " $$3; if ($$3+0 < 49.9) { print "Insufficient coverage"; exit 1; } }'

