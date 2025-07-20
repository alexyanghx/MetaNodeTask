pragma solidity ^0.8.20;

import "@chainlink/local/src/data-feeds/MockV3Aggregator.sol";
contract EthMockV3Aggregator is MockV3Aggregator {
    constructor(uint8 decimals,int256 price) 
        MockV3Aggregator(decimals, price) {
    }
}