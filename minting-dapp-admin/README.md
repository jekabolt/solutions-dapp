# <minting-dapp> admin panel

### tech stack
`react`
`webpack`
`typescript`
`SCSS`

### env
`node version >= 14`

### run locally
to clone and install deps:
```
git clone git@github.com:jekabolt/solutions-dapp.git
cd ./solutions-dapp/minting-dapp-admin
yarn install
```
to run dev server:
`yarn dev`

to build:
`yarn build`

todo:
- [x] setup react app (typescript + react + webpack + scss)
- [ ] use [zod](https://zod.dev/) for types ()
- [ ] use [useQuery](https://tanstack.com/query/v4/docs/reference/useQuery?from=reactQueryV3&original=https://react-query-v3.tanstack.com/reference/useQuery) for fetching data (if is has okay bundle size...didnt check yet)
- [ ] change build dir to art-admin/..../static
- [ ] add docker + make files, add run commands to readme
- [ ] think of todos