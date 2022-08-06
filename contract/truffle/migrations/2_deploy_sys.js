const SYSToken = artifacts.require("SYSToken");

module.exports = function(deployer) {
    deployer.deploy(SYSToken, "SYSToken", "SYS", "ipfs://QmNwL29Ng2sREGnWBVBwK9mLAYEkMkJXrLpAg1zv2iBhvo/metadata.json", "ipfs://QmNwL29Ng2sREGnWBVBwK9mLAYEkMkJXrLpAg1zv2iBhvo/hidden.json");
};