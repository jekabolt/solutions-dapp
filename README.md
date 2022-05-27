truffle console 

module.exports = function(deployer) {
    deployer.deploy(myContract, "GRBPWR Token", "MNT", { "type": "BigNumber", "hex": "0xb1a2bc2ec50000" }, 10000, 1, "ipfs://QmNwL29Ng2sREGnWBVBwK9mLAYEkMkJXrLpAg1zv2iBhvo/metadata.json");
};