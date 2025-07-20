const { task } = require("hardhat/config")
const { networkConfig, developmentChains } = require("../helper_hardhat_config");

task("lock-and-cross")
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

       const destNFTPoolDeployment = await deployments.get("DestNFTPool");
       receiver = (await hre.ethers.getContractAt("DestNFTPool",destNFTPoolDeployment.address)).target;

       await simulator.requestLinkFromFaucet(deployer,ethers.parseEther("100"));

    }else{
        if(!destChainId){
                destChainId = networkConfig[network.config.chainId].companionChainSelector
        }

        if(!receiver){
            const destNFTPoolDeployment = await hre.companionNetworks["destChain"].deployments.get("DestNFTPool");
            receiver = (await hre.ethers.getContractAt("DestNFTPool",destNFTPoolDeployment.address)).target;
        }

        linkTokenAddress = networkConfig[network.config.chainId].linkToken;
    }

   
    console.log(`destChainId :${destChainId}`);
    console.log(`destNFTPool adddress:${receiver}`);
    

    const myNFTDeployment = await deployments.get("MyNFT");
    const myNFT = await hre.ethers.getContractAt("MyNFT",myNFTDeployment.address);
  

    const sourceNFTPoolDeployment =  await deployments.get("SourceNFTPool");
    const sourceNFTPool = await hre.ethers.getContractAt("SourceNFTPool",sourceNFTPoolDeployment.address);
      await myNFT.approve(sourceNFTPool.target,tokenId);


    const linkToken = await hre.ethers.getContractAt("LinkToken",linkTokenAddress);
    console.log(`sourceNFTPool before balance :${await linkToken.balanceOf(sourceNFTPool.target)}`)
//    const linkTokenTx = await linkToken.transfer(sourceNFTPool.target,hre.ethers.parseEther("10"));
//    if(developmentChains.includes(hre.network.name)){
//         await linkTokenTx.wait();
//    }else{
//          await linkTokenTx.wait(1);
//    }
    
    console.log(`sourceNFTPool after balance :${await linkToken.balanceOf(sourceNFTPool.target)}`)

    const tx = await sourceNFTPool.lockAndSendNFT(deployer,tokenId,receiver,destChainId);
    console.log(`lock and send nft tx:${tx.hash}`)
}); 

module.exports={}