const { task } = require("hardhat/config")
const { networkConfig, developmentChains } = require("../helper_hardhat_config");

task("burn-and-cross")
.addParam("tokenid", "tokenId to lock")
.addOptionalParam("chainselector","destination chainId")
.addOptionalParam("receiver","destination address")
.setAction(async function (taskArgs,hre) { 
    const {deployments,getNamedAccounts} = hre;
    const {deployer} = await getNamedAccounts();

    const tokenId = taskArgs.tokenid;
    console.log(`tokenId :${tokenId}`);

    let destChainId = taskArgs.chainselector;
    let receiver = taskArgs.receiver; 

    let linkTokenAddress;
    if(developmentChains.includes(network.name)){
        const simulatorDeployment = await deployments.get("CCIPLocalSimulator");
        const simulator = await hre.ethers.getContractAt("CCIPLocalSimulator",simulatorDeployment.address);
        ccipLocalConfig = await simulator.configuration();
        destChainId = ccipLocalConfig.chainSelector_;
        linkTokenAddress = ccipLocalConfig.linkToken_;

       const destNFTPoolDeployment = await deployments.get("SourceNFTPool");
       receiver = (await hre.ethers.getContractAt("SourceNFTPool",destNFTPoolDeployment.address)).target;

       await simulator.requestLinkFromFaucet(deployer,ethers.parseEther("100"));

    }else{
        if(!destChainId){
                destChainId = networkConfig[network.config.chainId].companionChainSelector
        }

        if(!receiver){
            const sourceNFTPoolDeployment = await hre.companionNetworks["sourcechain"].deployments.get("SourceNFTPool");
            receiver = await hre.ethers.getContractAt("SourceNFTPool",sourceNFTPoolDeployment.address).target;
        }

        linkTokenAddress = networkConfig[network.config.chainId].linkToken;
    }

   
    console.log(`destChainId :${destChainId}`);
    console.log(`destNFTPool adddress:${receiver}`);
    

    const wmyNFTDeployment = await deployments.get("WrapperMyNFT");
    const wmyNFT = await hre.ethers.getContractAt("WrapperMyNFT",wmyNFTDeployment.address);
  

    const destNFTPoolDeployment =  await deployments.get("DestNFTPool");
    const destNFTPool = await hre.ethers.getContractAt("DestNFTPool",destNFTPoolDeployment.address);
    await wmyNFT.approve(destNFTPool.target,tokenId);


    const linkToken = await hre.ethers.getContractAt("LinkToken",linkTokenAddress);
    console.log(`sourceNFTPool before balance :${await linkToken.balanceOf(destNFTPool.target)}`)
   const linkTokenTx = await linkToken.transfer(destNFTPool.target,hre.ethers.parseEther("10"));
   if(developmentChains.includes(hre.network.name)){
        await linkTokenTx.wait();
   }else{
         await linkTokenTx.wait(6);
   }
    
    console.log(`destNFTPool after balance :${await linkToken.balanceOf(destNFTPool.target)}`)

    const tx = await destNFTPool.burnAndSendNFT(deployer,tokenId,receiver,destChainId);
    console.log(`burn and send nft tx:${tx.hash}`)
}); 

module.exports={}