const SYSToken = artifacts.require("SYSToken");

contract('SYSToken', () => {

    it('check hidden', async() => {
        const sys = await SYSToken.deployed();

        await sys.mint(1)
        let tokenUri = (
            await sys.tokenURI.call(1)
        )

        const hiddenUri = (
            await sys.notRevealedUri.call()
        )

        assert.equal(tokenUri, hiddenUri, "token uri should be hidden before reveal");

        await sys.reveal()

        testBaseUri = "ipfs://test/"

        await sys.setBaseURI(testBaseUri)

        tokenUri = (
            await sys.tokenURI.call(1)
        )

        let baseExtension = (
            await sys.baseExtension.call()
        )

        assert.equal(tokenUri, testBaseUri + "1" + baseExtension, "token uri should equal to testBaseUri");
    });

    it('check mint on pause', async() => {
        const sys = await SYSToken.deployed();
        await sys.pause(true)

        try {
            expect((
                await sys.mint(1)
            )).to.be.an('error');
        } catch (err) {
            expect(err.hijackedStack.includes("revert"), 'check mint on pause').to.be.true
        }
        await sys.pause(false)

    });
    it('check ownership', async() => {
        const sys = await SYSToken.deployed();

        let owner = (
            await sys.ownerById(1)
        )

        try {
            expect((
                await sys.ownerById(2)
            )).to.be.an('error');
        } catch (err) {
            expect(err.hijackedStack.includes("invalid token ID"), 'token with id 2 should be not minted yet').to.be.true
        }

        await sys.mint(1)
        let owner2 = (
            await sys.ownerById(2)
        )
        assert.equal(owner, owner2, "token 2 should be minted");


        let ownerAssets = (await sys.walletOfOwner(owner))
        expect(ownerAssets.length, 'minted twice should be equal 2').to.equal(2)
    });

    it('check max min amount', async() => {
        const sys = await SYSToken.deployed();
        let maxMintAmount = (await sys.maxMintAmount.call())
        expect(maxMintAmount.words[0], 'default max mint amount should be equal 5').to.equal(5)

        await sys.setmaxMintAmount(6)

        maxMintAmountAfter = (await sys.maxMintAmount.call())

        expect(maxMintAmountAfter.words[0], 'max mint amount should be changed to 6').to.equal(6)

        try {
            expect((
                await sys.mint(7)
            )).to.be.an('error');
        } catch (err) {
            expect(err.hijackedStack.includes("revert"), 'cannot mint more than max amount').to.be.true
        }
    });



});