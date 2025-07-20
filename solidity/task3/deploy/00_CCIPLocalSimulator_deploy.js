


module.exports=async (hre)=>{
    const {deployments,getNamedAccounts} =hre;
    const {deployer} = await getNamedAccounts();
    const {log} =deployments;

    log("deploying CCIPLocalSimulator on local network...")
    await deployments.deploy("CCIPLocalSimulator",{
        from: deployer,
        contract: "CCIPLocalSimulator",
        args:[],
        log: true
    })
    log("deployed CCIPLocalSimulator")
}

module.exports.tags=["mock","all"]