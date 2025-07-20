const {developmentChains,networkConfig} = require("../helper_hardhat_config");

module.exports = async (hre) => { 
    const {deployments,getNamedAccounts}= hre;
    const {log} = deployments;
    const {deployer} = await getNamedAccounts();

    let router,linkToken;
    log(`network name=${network.name}`)
    log(`in local :${developmentChains.includes(network.name)}`);
    if(developmentChains.includes(network.name)){
        const ccipLocalDeployment = await deployments.get("CCIPLocalSimulator");
        const CCIPLocalSimulator = await hre.ethers.getContractAt("CCIPLocalSimulator",ccipLocalDeployment.address);
        const config = await CCIPLocalSimulator.configuration();
        router = config.sourceRouter_;
        linkToken = config.linkToken_;
        log(`network.name=${network.name},CCIPLocalSimulator router=${router},linkToken=${linkToken}`)
    }else{
        const config = networkConfig[network.config.chainId];
        router = config.router;
        linkToken = config.linkToken;

        log(`network.name=${network.name},router=${router},linkToken=${linkToken}`)
    }
   

    const myNFTDeployment = await deployments.get("MyNFT");
    const nftAddr = myNFTDeployment.address;

   log(`deploying SourceNFTPool on source chain`);

   await deployments.deploy("SourceNFTPool",{
        from: deployer,
        contract:"SourceNFTPool",
        args:[router,linkToken,nftAddr],
        log: true
    })

    log(`deployed SourceNFTPool`);

};

module.exports.tags = ["sourcechain","all"];