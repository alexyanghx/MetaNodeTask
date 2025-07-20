
const {tokens} = require("../helper_hardhat_config.js");

module.exports=async (hre)=>{ 
    const {deployments,getNamedAccounts} =hre;
    const {deployer} = await getNamedAccounts();
    const{log} = deployments;

    const deplomentInfo = await deployments.deploy("UsdcMock",{
        contract:"UsdcMock",
        from:deployer,
        args:[],
        log:true
    })
  
    tokens.usdcTokenAddress=deplomentInfo.address;   
    log(`usdcTokenAddress:${tokens.usdcTokenAddress}`)
}

module.exports.tags=["mock"]