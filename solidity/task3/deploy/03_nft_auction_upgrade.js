module.exports=async (hre)=>{
    const {deployments,getNamedAccounts}=hre;
    const {deployer}=await getNamedAccounts();

    await deployments.fixture(["AuctionV1"]);

    const proxy = await deployments.get("auctionNFT");
    const factory = await hre.ethers.getContractFactory("AuctionNFT")
    const contract = await hre.upgrades.upgradeProxy(proxy.address,factory);
    await contract.waitForDeployment();

    console.log("contract proxy address:",contract.target)
    console.log("contract implementation address:",await hre.upgrades.erc1967.getImplementationAddress(contract.target))

    await deployments.save("AuctionV2",{
        address: contract.target,
        abi:factory.interface.format("json"),
        from:deployer,
        contractName:"AuctionNFT",
        args:[],
        log:true
    })
}



module.exports.tags = ["auctionV2"];