const SYSToken = artifacts.require("SYSToken");

contract('SYSToken', () => {
    it('initializes the contract with the correct values', function() {
        return SYSToken.deployed().then(function(instance) {
            sysInstance = instance;
            return sysInstance.address
        }).then(function(address) {
            assert.notEqual(address, 0x0, 'has contract address');
            return sysInstance.mint(1);
        }).then(function(address) {
            sysInstance.reveal.call()
        }).then(function() {
            return sysInstance.tokenURI(1)
        }).then(function(tokenUri) {
            console.log("-----", tokenUri);
        });
    });
    it('check hidden', async() => {
        // const sys = await SYSToken.deployed();

        // await sys.mint.call(1).then(function(sys) {
        //     const tokenUri = (
        //         await sys.tokenURI.call(1)
        //     )
        //     console.log("----- ", tokenUri);
        // })


    });
});