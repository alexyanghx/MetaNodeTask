const developmentChains = ["hardhat", "localhost"];

const tokens={
    ethTokenAddress: "0x0000000000000000000000000000000000000000",
    usdcTokenAddress: "0x1c7d4b196cb0c7b01d743fbc6116a902379c7238"
}

const mockFeedPrices ={
        [tokens.ethTokenAddress]:
        {
            decimals:8,
            price: 361147650000
        },
        [tokens.usdcTokenAddress]:
        {
            decimals: 8,
            price: 99985934
        }
}

const networkConfig = {
    11155111: {
        name: "sepolia",
        router: "0x0BF3dE8c5D3e8A2B34D2BEeB17ABfCeBaf363A59",
        linkToken: "0x779877A7B0D9E8603169DdbD7836e478b4624789",
        companionChainSelector: "16281711391670634445",
        usdcTokenAddress:"0x1c7d4b196cb0c7b01d743fbc6116a902379c7238",
        priceFeeds: 
        {
            [tokens.ethTokenAddress]:"0x694AA1769357215DE4FAC081bf1f309aDC325306",
            "0x1c7d4b196cb0c7b01d743fbc6116a902379c7238":"0xA2F78ab2355fe2f984D808B5CeE7FD0A93D5270E"
        }
    },
    80002: {
        name: "amoy",
        router: "0x9C32fCB86BF0f4a1A8921a9Fe46de3198bb884B2",
        linkToken: "0x0Fd9e8d3aF1aaee056EB9e802c3A762a667b1904",
        companionChainSelector: "16015286601757825753"
    }

}

module.exports ={
    developmentChains,
    networkConfig,
    tokens,
    mockFeedPrices
}