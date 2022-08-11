install: ## Install the web dependencies
	yarn install --ignore-engines

dev: ## Run the local dev server
	yarn dev

build-dist: ## Build dist version
	yarn build

build-proto: ## Generate proto-http files from nft.proto and auth.proto
	protoc \
		--typescript-http_out=./src/api/proto-http \
		--proto_path=../art-admin/proto \
		../art-admin/proto/nft/nft.proto ../art-admin/proto/auth/auth.proto 
