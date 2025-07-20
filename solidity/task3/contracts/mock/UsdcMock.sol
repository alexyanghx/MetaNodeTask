pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract UsdcMock is ERC20{
    constructor() ERC20("USD Coin", "USDC") {
        _mint(msg.sender, 1000*(10**decimals()));
    }


    /**
     * @dev Returns the decimals places of the token.
     */
    function decimals() public override pure returns (uint8){
        return 8;
    }
}