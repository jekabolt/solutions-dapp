// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "../contracts/SYSToken.sol";
// These files are dynamically created at test time
import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";

contract SYSTokenTest {

  function testWriteValue() public {
    SYSToken sys = SYSToken(DeployedAddresses.SYSToken());
    sys.mint(1);
    
    // Assert.equal(sys.tokenURI(1), sys.notRevealedUri, "Contract should have 1 stored");
  }
}
