const{developmentChains,networkConfig} = require("../helper_hardhat_config");

module.exports = async (hre) => { 
    const {deployments,getNamedAccounts}= hre;
    const {deployer,user1} = await getNamedAccounts();
    const {log} = deployments;

    let router,linkToken;
    if(developmentChains.includes(network.name)){

        const ccipLocalDeployment = await deployments.get("CCIPLocalSimulator");
        const CCIPLocalSimulator = await hre.ethers.getContractAt("CCIPLocalSimulator",ccipLocalDeployment.address);
        const config = await CCIPLocalSimulator.configuration();
        router = config.destinationRouter_;
        linkToken = config.linkToken_;

        log(`dest chain network name=${network.name},CCIPLocalSimulator router=${router},linkToken=${linkToken}`)
    }else{
        const config = networkConfig[network.config.chainId];
        router = config.router;
        linkToken = config.linkToken;
        
        log(`dest chain network name=${network.name},router=${router},linkToken=${linkToken}`)

    }


    const wmyNFTDeployment = await deployments.get("WrapperMyNFT");
    const wnftAddr = wmyNFTDeployment.address;

    log(`deploying DestNFTPool on dest chain....`);

   await deployments.deploy("DestNFTPool",{
        from: deployer,
        contract:"DestNFTPool",
        args:[router,linkToken,wnftAddr],
        log: true
    })
  log(`deployed DestNFTPool`);
};

module.exports.tags = ["destchain","all"];