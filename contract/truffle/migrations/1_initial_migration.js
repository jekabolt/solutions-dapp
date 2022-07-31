/* eslint-disable prettier/prettier */
/* eslint-disable no-undef */
const Migrations = artifacts.require("Migrations");

module.exports = function(deployer) {
    deployer.deploy(Migrations);
};