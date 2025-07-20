
const {task} = require("hardhat/config");

task("check-wrapper-nft")
.addParam("tokenid","token id")
.setAction(async function (taskArgs,hre){ 
    const {deployments,getNamedAccounts} = hre;
    const tokenId = taskArgs.tokenid;

    const wrapperMyNFTDeployment = await deployments.get("WrapperMyNFT");

    const wnft = await hre.ethers.getContractAt("WrapperMyNFT",wrapperMyNFTDeployment.address);
 console.log("2")
    const totalSupply = await wnft.totalSupply();
    console.log(`totalSupply:${totalSupply}`);
    console.log(`tokenId:${tokenId} owner:${await wnft.ownerOf(tokenId)}`);
});

module.exports= {
}