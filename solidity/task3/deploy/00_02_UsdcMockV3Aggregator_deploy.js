
const {tokens,mockFeedPrices} = require("../helper_hardhat_config.js");

module.exports=async (hre)=>{ 
    const {deployments,getNamedAccounts} =hre;
    const {deployer} = await getNamedAccounts();

    const price = mockFeedPrices[tokens.usdcTokenAddress]["price"];
    const decimals = mockFeedPrices[tokens.usdcTokenAddress]["decimals"];
    
    await deployments.deploy("UsdcMockV3Aggregator",{
        contract:"UsdcMockV3Aggregator",
        from:deployer,
        args:[decimals,price],
        log:true
    })
}

module.exports.tags=["auction","mock"]