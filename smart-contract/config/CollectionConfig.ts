/* eslint-disable node/no-missing-import */
import CollectionConfigInterface from "../lib/CollectionConfigInterface";
import {
  ethereumTestnet,
  ethereumMainnet,
  hardhatLocal,
} from "../lib/Networks";
import { openSea } from "../lib/Marketplaces";
import whitelistAddresses from "./whitelist.json";

const CollectionConfig: CollectionConfigInterface = {
  testnet: ethereumTestnet,
  mainnet: ethereumMainnet,
  hardhat: hardhatLocal,
  // The contract name can be updated using the following command:
  // yarn rename-contract NEW_CONTRACT_NAME
  // Please DO NOT change it manually!
  contractName: "SOUTIONSToken",
  tokenName: "GRBPWR Token",
  tokenSymbol: "MNT",
  hiddenMetadataUri: "ipfs://__CID__/hidden.json",
  maxSupply: 10000,
  whitelistSale: {
    price: 0.05,
    maxMintAmountPerTx: 1,
  },
  preSale: {
    price: 0.07,
    maxMintAmountPerTx: 2,
  },
  publicSale: {
    price: 0.09,
    maxMintAmountPerTx: 5,
  },
  contractAddress: "0x369a2ff91eEB5a440E57416aB479467Cb6d3Cf1b",
  marketplaceIdentifier: "my-nft-token",
  marketplaceConfig: openSea,
  whitelistAddresses: whitelistAddresses,
};

export default CollectionConfig;
