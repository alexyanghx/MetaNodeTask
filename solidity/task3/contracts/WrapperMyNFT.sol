//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./MyNFT.sol";

contract WrapperMyNFT is MyNFT{
    uint256 private _tokenIdCounter;
    mapping(uint256 => string) private _tokenURIs;

    constructor(string memory name,string memory symbol) MyNFT(name, symbol){}

   function safeMint(address to, uint256 tokenId) public{
      super._safeMint(to, tokenId);
   }

   function burn(uint256 tokenId) public{
      super._burn(tokenId);
   }

}