module.exports = async (hre) => { 
    const { deployments, getNamedAccounts } = hre;
    const { deployer } = await getNamedAccounts();
    const {log} = deployments;
    // const factory = await hre.ethers.getContractFactory("AuctionNFT");

    // const proxy = await hre.upgrades.deployProxy(factory,[],{"initializer": "initialize"});
    // await proxy.waitForDeployment();

    // console.log("contract proxy address:",proxy.target)
    // console.log("contract implementation address:",await hre.upgrades.erc1967.getImplementationAddress(proxy.target))

    // await deployments.save("auctionNFT",{
    //     address: proxy.target,
    //     abi: proxy.interface.format("json"),
    //     from: deployer,
    //     contractName:"AuctionNFT",
    //     args:[],
    //     log:true
    // })
    log(`deploying AuctionNFT...`)
    await deployments.deploy("AuctionNFT",{
        contract:"AuctionNFT",
        from:deployer,
        proxy:{
            proxyContract:"OpenZeppelinTransparentProxy",
            ViaAdminContract:"DefaultProxyAdmin",
            execute:{
                methodName:"initialize",
                args:[deployer]
            }
        },
        log:true,
        args:[]
    })
    log(`deployed AuctionNFT`)
};

module.exports.tags = ["auction"]