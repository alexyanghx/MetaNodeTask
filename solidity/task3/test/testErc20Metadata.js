
describe("erc20Metadata test",function(){ 

    it("run test",async function(){
        const contractAddress="0x1c7d4b196cb0c7b01d743fbc6116a902379c7238";
        // 方案1：使用OpenZeppelin官方路径
const token = await ethers.getContractAt(
  "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol:IERC20Metadata", 
  contractAddress
);

// // 方案2：使用本地依赖路径
// const token = await ethers.getContractAt(
//   "contracts/.deps/npm/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol:IERC20Metadata",
//   contractAddress
// );

        // const token = await ethers.getContractAt("IERC20Metadata",contractAddress);
        console.log(await token.decimals());
    })
    
})