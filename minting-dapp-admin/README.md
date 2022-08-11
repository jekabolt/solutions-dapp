# <minting-dapp> admin panel

### env
`node version >= 16`

### clone
```
git clone git@github.com:jekabolt/solutions-dapp.git
git checkout dev-admin
cd ./solutions-dapp/minting-dapp-admin
```

### run locally without Makefile
```
yarn install #to install

yarn dev #to run

yarn build #to build

yarn build-proto #to generate proto-http
```

### run locally without Makefile
```
make install #to install

make dev #to run

make build #to build

make build-proto #to generate proto-http
```

to use `make generate-proto` make sure you have GO `.bin`, `protoc` compiler and `proto->http` ([protoc-gen-typescript-http](https://github.com/einride/protoc-gen-typescript-http)) generator  available in working dir 

todo:
- [x] setup react app (typescript + react + webpack + scss)
- [ ] generate types from backend (protobuff,rpc) (similar to [tRPC](https://trpc.io/docs/) types generation ?))
- [ ] use [useQuery](https://tanstack.com/query/v4/docs/reference/useQuery?from=reactQueryV3&original=https://react-query-v3.tanstack.com/reference/useQuery) for fetching data (if is has okay bundle size...didnt check yet)
- [ ] change build dir to art-admin/..../static
- [ ] add docker + make files, add run commands to readme
- [ ] color mode (theme)
- [ ] check bundle routing
- [ ] prior todos
- [ ] add protected wrapper to routes with auth
- [ ] move proto files to upper scope (both client and server generates it)
- [ ] think of todos