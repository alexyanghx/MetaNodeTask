// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721, Ownable{

    uint private _tokenIdCounter;
    mapping(uint => string) private _tokenURIs;

    constructor() ERC721("MyNFT", "MNFT") Ownable(msg.sender) {}
    function mintNFT(address to,string memory uri) public onlyOwner {
        _tokenIdCounter++;
        uint tokenId = _tokenIdCounter;
        _mint(to, tokenId);
        _tokenURIs[tokenId] = uri;
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory)  {
        return _tokenURIs[tokenId];
    }
}