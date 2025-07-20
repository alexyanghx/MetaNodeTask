
module.exports=async(hre)=>{
    const {deployments,getNamedAccounts,ethers} =hre;
    const{deployer} = await getNamedAccounts();
    const {log} =deployments;

    log("deploying MyNFT on source chain...")
    await deployments.deploy("MyNFT",{
        contract: "MyNFT",
        from: deployer,
        args:["MyNFT","MT"],
        log: true
    })
    log("deployed MyNFT");
}

module.exports.tags = ["sourcechain","auction","all"]