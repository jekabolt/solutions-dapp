# <minting-dapp> admin panel

<<<<<<< HEAD
### tech stack
=======
### core tech stack
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
`react`
`webpack`
`typescript`
`SCSS`

### env
<<<<<<< HEAD
`node version >= 14`
=======
`node version >= 15.0.1` (make sure its correct, im using 16.16)
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c

### run locally
to clone and install deps:
```
git clone git@github.com:jekabolt/solutions-dapp.git
cd ./solutions-dapp/minting-dapp-admin
<<<<<<< HEAD
=======
git checkout dev-admin
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
yarn install
```
to run dev server:
`yarn dev`

to build:
`yarn build`

todo:
- [x] setup react app (typescript + react + webpack + scss)
<<<<<<< HEAD
- [ ] use [zod](https://zod.dev/) for types ()
=======
- [ ] generate types from backend (protobuff,rpc) (similar to [tRPC](https://trpc.io/docs/) types generation ?))
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
- [ ] use [useQuery](https://tanstack.com/query/v4/docs/reference/useQuery?from=reactQueryV3&original=https://react-query-v3.tanstack.com/reference/useQuery) for fetching data (if is has okay bundle size...didnt check yet)
- [ ] change build dir to art-admin/..../static
- [ ] add docker + make files, add run commands to readme
- [ ] color mode (theme)
<<<<<<< HEAD
=======
- [ ] check bundle routing
- [ ] prior todos
- [ ] add protected wrapper to routes with auth
- [ ] move proto files to upper scope (both client and server generates it)
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
- [ ] think of todos