require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require("@chainlink/env-enc").config()
require("@nomicfoundation/hardhat-ethers");
require("hardhat-deploy");
require("hardhat-deploy-ethers");
require("./task")

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks:{
    sepolia:{
      chainId: 11155111,
      url: process.env.SEPOLIA_RPC_URL,
      accounts: [process.env.AC_PK1,process.env.AC_PK2],
      blockConfirmations: 6,
      companionNetworks:{
        destChain: "amoy"
      }
    },
    amoy:{
      chainId: 80002,
      url: process.env.AMOY_RPC_URL,
      accounts: [process.env.AC_PK1,process.env.AC_PK2],
      blockConfirmations: 6,
      timeout: 1000000,
      companionNetworks:{
        destChain: "sepolia"
      }
    },
    localhost:{
      url: "http://127.0.0.1:8545/"
    }
  },
  namedAccounts:{
    deployer:{
      default:0
    },
    user1:{
      default:1
    },
    user2:{
      default:2
    }
  }
};
