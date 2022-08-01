# <minting-dapp> admin panel

### core tech stack
`react`
`webpack`
`typescript`
`SCSS`

### env
`node version >= 15.0.1` (make sure its correct, im using 16.16)

### run locally
to clone and install deps:
```
git clone git@github.com:jekabolt/solutions-dapp.git
cd ./solutions-dapp/minting-dapp-admin
git checkout dev-admin
yarn install
```
to run dev server:
`yarn dev`

to build:
`yarn build`

todo:
- [x] setup react app (typescript + react + webpack + scss)
- [ ] generate types from backend (protobuff,rpc) (similar to [tRPC](https://trpc.io/docs/) types generation ?))
- [ ] use [useQuery](https://tanstack.com/query/v4/docs/reference/useQuery?from=reactQueryV3&original=https://react-query-v3.tanstack.com/reference/useQuery) for fetching data (if is has okay bundle size...didnt check yet)
- [ ] change build dir to art-admin/..../static
- [ ] add docker + make files, add run commands to readme
- [ ] color mode (theme)
- [ ] check bundle routing
- [ ] think of todos