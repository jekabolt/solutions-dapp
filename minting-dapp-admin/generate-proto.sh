#!/bin/bash
 
# generate js codes via grpc-tools
# yarn run grpc_tools_node_protoc \
#   --js_out=import_style=commonjs,binary:./src/proto \
#   --grpc_out=./src/proto \
#   --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin \
#   -I ../art-admin/proto \
#   ../art-admin/proto/auth/auth.proto ../art-admin/proto/nft/nft.proto

# # generate d.ts codes
# yarn run grpc_tools_node_protoc \
#   --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
#   --ts_out=./src/proto \
#   -I ../art-admin/proto \
#   ../art-admin/proto/auth/auth.proto ../art-admin/proto/nft/nft.proto

yarn run grpc_tools_node_protoc --ts_out ./src/proto --proto_path ../art-admin/proto ../art-admin/proto/nft/nft.proto ../art-admin/proto/auth/auth.proto

echo 'you cool bro'

# npx protoc --ts_out . --proto_path protos protos/msg-readme.proto

