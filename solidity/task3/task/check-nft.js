task("check-nft").setAction(async function (taskArgs,hre){
    const {deployments,getNamedAccounts} = hre;

   const myNFTDeployment =await deployments.get("MyNFT");

   const myNFT = await hre.ethers.getContractAt("MyNFT",myNFTDeployment.address);

    const totalSupply = await myNFT.totalSupply();
   for (let i = 1; i <= totalSupply; i++){
        console.log(`tokenId:${i} owner:${await myNFT.ownerOf(i)}`);
   }
})