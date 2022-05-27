/* eslint-disable no-undef */
/* eslint-disable prettier/prettier */
// import CollectionConfig from "../config/CollectionConfig";
// import { NftContractType } from "../lib/NftContractProvider";

const myContract = artifacts.require("SOUTIONSToken");

module.exports = function(deployer) {
    deployer.deploy(myContract, "SOUTIONSToken", "SOUTION", "ipfs://QmNwL29Ng2sREGnWBVBwK9mLAYEkMkJXrLpAg1zv2iBhvo/metadata.json", "ipfs://QmNwL29Ng2sREGnWBVBwK9mLAYEkMkJXrLpAg1zv2iBhvo/hidden.json");
};