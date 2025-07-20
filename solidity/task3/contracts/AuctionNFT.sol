// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "hardhat/console.sol";

contract AuctionNFT is IERC721Receiver,Initializable{
    struct Auction {
        uint256 id;
        uint256 minimumBid;
        uint256 duration;
        uint256 startTime;
        bool ended;

        address nftAddress;
        uint256 tokenId;

        address seller;
        address highestBidder;
        uint256 highestBid;
        address tokenAddress;
    }

    address private _admin;
    uint private _auctionIdCounter;
    mapping(uint256 => Auction) private _auctions;
    mapping(address=>AggregatorV3Interface) private _priceFeeds;

    event AuctionCreated(
            uint256 indexed auctionId,
            address indexed nftAddress,
            uint256 tokenId,
            address seller,
            uint256 startTime,
            uint256 duration,
            uint256 minimumBid
        );

    modifier onlyOwner(uint auctionId) {
        require(msg.sender == _admin||msg.sender==_auctions[auctionId].seller, "Only admin or seller can call this function.");
        _;
    }

    function initialize(address router) public initializer{
        _admin = msg.sender;
    }

    function createAuction(
        uint256 _auctionDuration,
        uint256 _minimumBid,
        address _nftAddress,
        uint256 _tokenId
    ) public { 

        require(
           _auctionDuration>10,
            "_auctionDuration must be greater than 10."
        );

        require(_minimumBid>0,"_minimumBid must be greater than 0.");

        require(_nftAddress!=address(0), "nftAddress cannot be 0.");
        
        IERC721 erc721 = IERC721(_nftAddress);
        require(erc721.ownerOf(_tokenId) == msg.sender, "You are not the owner of this NFT.");


        Auction memory newAuction = Auction({
            id: _auctionIdCounter,
            startTime: block.timestamp,
            duration: _auctionDuration,
            minimumBid: _minimumBid,
            highestBid: 0,
            highestBidder: address(0),
            ended: false,
            seller:msg.sender,
            tokenId: _tokenId,
            nftAddress: _nftAddress,
            tokenAddress:address(0)
            });
       
        _auctions[_auctionIdCounter] = newAuction;
        _auctionIdCounter++;

        emit AuctionCreated(newAuction.id,newAuction.nftAddress,newAuction.tokenId,newAuction.seller,newAuction.startTime,newAuction.duration,newAuction.minimumBid);
    }

    function bid(uint256 _auctionId,uint amount,address tokenAddress) public payable {
       
       Auction storage auction = _auctions[_auctionId];

       require(!auction.ended&&block.timestamp < auction.duration + auction.startTime, "Auction has already ended.");

        uint payValue;
        //使用eth账户支付
       if(tokenAddress==address(0)){
            amount = msg.value;  
       }
    console.log(3);
        payValue = getUsdPrice(amount,tokenAddress);
            console.log(4);
        uint highestBidUsd = getUsdPrice(auction.highestBid,auction.tokenAddress);
            console.log(5);
        require(payValue > highestBidUsd, "Your bid is lower than the highest bid.");

        uint minimumBidUsd = getUsdPrice(auction.minimumBid,auction.tokenAddress);
        require(payValue > minimumBidUsd, "Your bid is lower than the minimum bid.");
      
      //存在竞价
      if(auction.highestBidder!=address(0)){
        if(address(0)==auction.tokenAddress){
            payable(auction.highestBidder).transfer(auction.highestBid);
        }else{
            IERC20(auction.tokenAddress).transfer(auction.highestBidder,auction.highestBid);
        }
      }


        auction.highestBidder = msg.sender;
        auction.highestBid = amount;
        auction.tokenAddress = tokenAddress;

    

    }

    function endAuction(uint256 _auctionId) public onlyOwner(_auctionId) {
        Auction storage auction = _auctions[_auctionId];
        require(!auction.ended, "Auction had ended.");
        require(auction.startTime+auction.duration <= block.timestamp, "Auction not ended yet");

        auction.ended = true;
        IERC721 erc721 = IERC721(auction.nftAddress);
        address owner = auction.seller;

        console.log(1);
        if (auction.highestBidder != address(0)) {
            
            erc721.safeTransferFrom(owner,auction.highestBidder, auction.tokenId);
            console.log(2);
            console.log(auction.tokenAddress);    
            if(auction.tokenAddress == address(0)){
                payable(owner).transfer(auction.highestBid);
            }else{
                console.log(3);
                IERC20 erc20 = IERC20(auction.tokenAddress);
                erc20.transfer(owner, auction.highestBid);
            }
        }
    }

    function getAuction(uint256 _auctionId) public view returns (Auction memory){
        require(_auctionId < _auctionIdCounter, "Invalid auction ID.");
        return _auctions[_auctionId];
    }



    function getUsdPrice(uint256 amount, address tokenAddress) private view returns (uint256) {

        // 获取当前价格
        uint256 price = uint256(getChainlinkDataFeedLatestAnswer(tokenAddress));

        if(address(0)==tokenAddress){
            return amount*price/1e18;
        }

        try IERC20Metadata(tokenAddress).decimals() returns (uint8 decimals) {

            price= amount*price/10**decimals;

        } catch {
    
            price = amount*price/1e18;

        }
        return price;
    }

     function getChainlinkDataFeedLatestAnswer(address tokenAddress) public view returns (int256 answer) {
        require(hadPriceFeed(tokenAddress),"no price feed");

        (,answer,,,) = _priceFeeds[tokenAddress].latestRoundData();
        return answer;
     }
       
       
    function setPriceFeed(address tokenAddress,address priceFeedAddress) public{
        require(msg.sender==_admin,"only admin can set price feed");

        _priceFeeds[tokenAddress]=AggregatorV3Interface(priceFeedAddress);
    }

    function hadPriceFeed(address tokenAddress) public view returns (bool){
        return address(_priceFeeds[tokenAddress])!=address(0);
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external pure returns (bytes4) { 
        return IERC721Receiver.onERC721Received.selector;
    }
}