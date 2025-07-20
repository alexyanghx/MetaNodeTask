
const {tokens,mockFeedPrices} = require("../helper_hardhat_config.js");

module.exports=async (hre)=>{ 
    const {deployments,getNamedAccounts} =hre;
    const {deployer} = await getNamedAccounts();

    const price = mockFeedPrices[tokens.ethTokenAddress]["price"];
    const decimals = mockFeedPrices[tokens.ethTokenAddress]["decimals"];
    
    await deployments.deploy("EthMockV3Aggregator",{
        contract:"EthMockV3Aggregator",
        from:deployer,
        args:[decimals,price],
        log:true
    })
}

module.exports.tags=["mock"]