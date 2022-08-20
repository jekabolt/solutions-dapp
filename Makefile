REGISTRY=jekabolt
VERSION=master

build: statics
	cd art-admin && go build -a $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o ./bin/$(IMAGE_NAME) ./cmd/	

run: build
	cd art-admin && ./bin/$(IMAGE_NAME)

test: generate
	# we only have non-generated code in ./internal, so we only count coverage for it
	cd art-admin && go test -cover -coverprofile coverage.out -coverpkg ./app/...,./internal/... ./...
	# IMPORTANT: required coverage can only be increased
	cd art-admin && go tool cover -func coverage.out | \
		awk 'END { print "Coverage: " $$3; if ($$3+0 < 40.0) { print "Insufficient coverage"; exit 1; } }'
		
local: build
	cd art-admin && source .env && ./bin/$(IMAGE_NAME)

generate:
	buf generate --path proto/nft/nft.proto \
	--path proto/auth/auth.proto

abi:
	solcjs --optimize --bin --abi --include-path contract/truffle/node_modules/ --base-path ./contract/truffle/contracts --output-dir ./bin  contract/truffle/contracts/SYSToken.sol
	solcjs --optimize --bin --abi --include-path contract/truffle/node_modules/ --base-path ./contract/truffle/contracts --output-dir ./bin  contract/truffle/contracts/SYSToken.sol

statics: generate 
	@echo "Create temp dir for static files"
	@mkdir -p art-admin/app/static/swagger/temp
	@echo "Generating combined Swagger JSON"
	@find art-admin/app/static/swagger -type f -name "*.json" -exec cp {} art-admin/app/static/swagger/temp \;
	@GOOS="" GOARCH="" go run art-admin/app/static/swagger/main.go art-admin/app/static/swagger/temp > art-admin/app/static/swagger/api.swagger.json
	@rm -rf art-admin/app/static/swagger/temp
	@echo "Generating combined Swagger JSON - Done"


clean:
	rm -rf art-admin/bin
	rm -f art-admin/app/static/swagger/auth/*.json
	rm -f art-admin/app/static/swagger/nft/*.json
	
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install golang.org/x/text/cmd/gotext@latest
	go install go.einride.tech/protoc-gen-typescript-http@latest


BUF_VERSION 		:= 1.7.0
BIN 				:= $(shell echo $(PATH))
OS 					:= $(shell uname -s)
ARCH 				:= $(shell uname -m)
buf-install:
	echo "Installing buf version $(BUF_VERSION)"
	BIN="/usr/local/bin" && \
	VERSION="1.7.0" && \
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(OS)-$(ARCH)" \
		-o "./bin/buf" && \
	chmod +x "bin/buf" 

# admin panel 

install-admin-panel: ## Install the web dependencies
	cd admin-panel && yarn install --ignore-engines

dev-admin-panel: ## Run the local dev server
	cd admin-panel && yarn dev

build-dist-admin-panel: ## Build dist version
	cd admin-panel && yarn build


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


