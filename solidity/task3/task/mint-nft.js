const {task} = require("hardhat/config")

task("mint-nft").addOptionalParam("owner","nft owner,default nft deployer").addParam("tokenuri","nft metadata uri").setAction(async function (taskArgs,hre) {
    
    
    const {deployments,getNamedAccounts} = hre;
    const {deployer} = await getNamedAccounts();

    let owner = taskArgs.owner;
    if(!owner){
        owner = deployer;
    }

    const tokenUri = taskArgs.tokenuri;

    console.log(`input owner:${owner},tokenuri:${tokenUri}`)

    const myNFTDeployment = await deployments.get("MyNFT");
    const myNFT = await hre.ethers.getContractAt("MyNFT",myNFTDeployment.address);
    // console.log(myNFT);
    const tx = await myNFT.mintNFT(owner,tokenUri);
    const b = await tx.wait();

    const filter = myNFT.filters.Transfer(null,owner);
    const events = await myNFT.queryFilter(filter,b.blockNumber,b.blockNumber);
    const tokenId = events[0].args.tokenId;
    const totalSupply = await myNFT.totalSupply();
    console.log(`totalSupply:${totalSupply},tokenId:${tokenId}`)
});

module.exports={}