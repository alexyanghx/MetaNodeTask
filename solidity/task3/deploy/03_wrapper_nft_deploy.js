
module.exports=async(hre)=>{
    const {deployments,getNamedAccounts,ethers} =hre;
    const{deployer} = await getNamedAccounts();
        const {log} = deployments;

    log("deploying WrapperMyNFT on destchain")
    await deployments.deploy("WrapperMyNFT",{
        contract: "WrapperMyNFT",
        from: deployer,
        args:["WrapperMyNFT","WMT"],
        log: true
    })
    log("deployed WrapperMyNFT");
}

module.exports.tags = ["destchain","all"]