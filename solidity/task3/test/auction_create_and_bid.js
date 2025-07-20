
const { ethers,deployments,getNamedAccounts, network } = require("hardhat");
const {expect} = require("chai");
const {
    developmentChains,
    tokens,
    networkConfig
} = require("../helper_hardhat_config");

// const {deployments,getNamedAccounts} = hre;

let nftContract,auctionContract;
let deployer,user1;
let usdcAggregatorContract,ethAggregatorContract;

const endDuration = 11;

before(async function(){
  
    const accounts = await getNamedAccounts();
    deployer = accounts.deployer;
    user1 = accounts.user1;

    await deployments.fixture(["auction"]);

    console.log("loading nft contract.......")
    const nftDeployment = await deployments.get("MyNFT");
    nftContract = await hre.ethers.getContractAt("MyNFT", nftDeployment.address); 
    console.log("load nft contract success")

    console.log("loading auction contract.......")
    const auctionDeployment = await deployments.get("AuctionNFT");
    auctionContract = await hre.ethers.getContractAt("AuctionNFT", auctionDeployment.address);
    console.log("load auction success")

    console.log("loading eth aggregator contract.......")
    const ethAggregatorDeployment = await deployments.get("EthMockV3Aggregator");
    ethAggregatorContract = await hre.ethers.getContractAt("EthMockV3Aggregator", ethAggregatorDeployment.address);
    console.log("load eth aggregator success")

    const usdcAggregatorDeployment = await deployments.get("UsdcMockV3Aggregator");
    usdcAggregatorContract = await hre.ethers.getContractAt("UsdcMockV3Aggregator", usdcAggregatorDeployment.address);
    console.log("load usdc aggregator success")
});


describe("test mint nft and create auction", async function(){
    it("mint nft and owner of deployer",async function(){
        await nftContract.mintNFT(deployer,"ipfs://bafkreiht5eqy3m3wjs4n67njkfysdytuh4e7q6ljjpiodz2nlbshdz2cz4");

        const owner = await nftContract.ownerOf(1);
        expect(owner).to.be.equal(deployer);
    });

    it("create auction",async function(){
        await auctionContract.createAuction(endDuration,ethers.parseEther("0.001"),nftContract.target,1);
        const auctionInfo = await auctionContract.getAuction(0);

        await nftContract.approve(auctionContract.target,1);

        expect(auctionInfo.nftAddress).to.be.equal(nftContract.target);
    });
});

describe("test set priceFeed", async function(){ 
    it("test set eth and usdc",async function(){ 

        if(developmentChains.includes(network.name)){
           const usdcTokenAddress = tokens.usdcTokenAddress;
           const ethTokenAddress = tokens.ethTokenAddress;

            await auctionContract.setPriceFeed(usdcTokenAddress,usdcAggregatorContract.target);
             console.log(`network local set usdctoken=[${usdcTokenAddress}],aggregator=[${usdcAggregatorContract.target}]`)

            await auctionContract.setPriceFeed(ethTokenAddress,ethAggregatorContract.target);
             console.log(`network local set ethtoken=[${ethTokenAddress}],aggregator=[${ethAggregatorContract.target}]`)
            console.log(1111)
            const usdcContract = await ethers.getContractAt("UsdcMock",usdcTokenAddress);
            console.log("deployer balanceOf ",await usdcContract.balanceOf(deployer));
            await usdcContract.connect(await ethers.getSigner(deployer)).transfer(auctionContract.target,10000000000);
            console.log(2222)
             const hadUsdcFeed = await auctionContract.hadPriceFeed(usdcTokenAddress);
            expect(hadUsdcFeed).to.be.true;

            const hadEthFeed = await auctionContract.hadPriceFeed(ethTokenAddress);
            expect(hadEthFeed).to.be.true;    
        }else{
            const priceFeeds = networkConfig[network.config.chainId].priceFeeds;
            for (key of priceFeeds){
                const aggregatorAddr = priceFeeds[key];
                await auctionContract.setPriceFeed(key,priceFeeds[key]);
                console.log(`network [${network.name}]set token=[${key}],aggregator=[${aggregatorAddr}]`)
                const hadEthFeed = await auctionContract.hadPriceFeed(key);
                expect(hadEthFeed).to.be.true;   
            }
        }
        
    })
});

describe("test user bid", async function(){
    it("test user1 bid use usdt",async function(){ 
        let usdcTokenAddress = tokens.usdcTokenAddress;
        if(!developmentChains.includes(network.name)){
            usdcTokenAddress = networkConfig[network.config.chainId].usdcTokenAddress;
        }
        const user = await hre.ethers.getSigner(user1);
        await auctionContract.connect(user).bid(0,1000000000,usdcTokenAddress);
        const auctionInfo = await auctionContract.getAuction(0);
        expect(auctionInfo.highestBidder).to.be.equal(user1);
    })
    
})

describe("test end auction", function(){
    it("test user1 end auction",async function(){ 
        await new Promise(async (resolve, reject) => { 
            setTimeout(resolve,endDuration*1000);
        });

        
        await auctionContract.connect(await hre.ethers.getSigner(deployer)).endAuction(0);
        const owner =await nftContract.ownerOf(1);
        expect(owner).to.be.equal(user1);
    })
})