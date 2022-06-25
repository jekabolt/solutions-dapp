# art-admin

is a backend service for uploading storing mint requests offchain and in chain 
basic flow of the backend:

1. Some user got paid for nft thru api on website after pictures and descriptions uploaded we put his mint request into pool with status `StatusUnknown`.
2. We have a watcher with is update statuses for txs if users mint request transaction was successful we put a `StatusPending` for this mint request or it also can be `StatusFailed` if tx is left from blockchain.
3. Than it can be appeared on admin panel where admin can review the pictures and create "NFT" from em.
4. "NFT" is done admin can upload them offchain this is made because we bulk the uploads than our mint request got status StatusUploadedOffchain`.
5. When we are ready to upload it to blockchain we can generate the metadata for our collection and than update the uri link for this in our smart contract



```
StatusUnknown          :"unknown"
StatusPending          :"pending"
StatusCompleted        :"completed"
StatusFailed           :"failed"
StatusBad              :"bad"
StatusUploadedOffchain :"uploadedOffchain"
StatusUploaded         :"uploaded"
```

the http api specification can be found [here](https://api.sys.solutions/nft/api)

## configuration 

| KEY                     | VALUE                                                               
|---------------------------------|--------------------------------------------------------------------|
| SERVER_PORT                     | 3999                                                               | 
| S3_ACCESS_KEY                   | xxx                                                                | 
| S3_SECRET_ACCESS_KEY            | xxx                                                                | 
| S3_ENDPOINT                     | fra1.digitaloceanspaces.com                                        | 
| S3_BUCKET_NAME                  | grbpwr                                                             | 
| S3_BUCKET_LOCATION              | fra-1                                                              | 
| S3_BASE_FOLDER                  | solutions                                                          | 
| S3_IMAGE_STORE_PREFIX           | grbpwr-com                                                         | 
| S3_METADATA_STORE_PREFIX        | metadata                                                           | 
| S3_IPFS_STORAGE_PATH            | ipfs                                                               | 
| BUNT_DB_PATH                    | /etc/bunt/storage.db                                               | 
| ETH_ETHERSCAN_API_KEY           | xxx                                                                | 
| ETH_ETHERSCAN_NETWORK           | api-rinkeby                                                        | 
| ETH_CONTRACT_ADDRESS            | 0x0TODO:                                                           | 
| ETH_REFRESH_TIME                | 10m                                                                | 
| MORALIS_API_KEY                 | xxx                                                                | 
| MORALIS_TIMEOUT                 | 10s                                                                | 
| MORALIS_BASE_URL                | https://deep-index.moralis.io/api/v2/                              | 
| DESCRIPTIONS_PATH               | etc/descriptions.json                                              | 
| DESCRIPTIONS_COLLECTION_NAME    | Solutions #                                                        | 
| NFT_TOTAL_SUPPLY                | 10000                                                              | 
| AUTH_JWT_SECRET                 | hehe                                                               | 
| AUTH_ADMIN_PASSWORD             | hehe                                                               | 
| AUTH_PASSWORD_HASHER_SALT       | 16                                                                 | 
| AUTH_PASSWORD_HASHER_ITERATIONS | 100000                                                             | 
| AUTH_JWT_TTL                    | 60m                                                                | 

### run locally 

```bash
make install
make run
```

### run in docker  

```bash
make image-build
make image-run
```

