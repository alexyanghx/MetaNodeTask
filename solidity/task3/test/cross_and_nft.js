
const { ethers,deployments,getNamedAccounts } = require("hardhat");
const {expect} = require("chai");

let nft,wnft,sourceNFTPool,destNFTPool,chainselector;
let deployer,user1;

before(async function(){
    const namedAccounts = await getNamedAccounts();
    deployer = namedAccounts.deployer;
    user1 = namedAccounts.user1;

    await deployments.fixture(["all"]);
    const nftDeployment = await deployments.get("MyNFT");
    nft = await ethers.getContractAt("MyNFT",nftDeployment.address);

    const wrapperMyNFTDeployment = await deployments.get("WrapperMyNFT");
    wnft = await ethers.getContractAt("WrapperMyNFT",wrapperMyNFTDeployment.address);

    const sourceNFTPoolDeployment = await deployments.get("SourceNFTPool");
    sourceNFTPool = await ethers.getContractAt("SourceNFTPool",sourceNFTPoolDeployment.address);


    const destNFTPoolDeployment = await deployments.get("DestNFTPool");
    destNFTPool = await ethers.getContractAt("DestNFTPool",destNFTPoolDeployment.address);

    const CCIPLocalSimulatorDeployment = await deployments.get("CCIPLocalSimulator");
    ccipLocalSimulator = await ethers.getContractAt("CCIPLocalSimulator",CCIPLocalSimulatorDeployment.address);
    chainselector = (await ccipLocalSimulator.configuration()).chainSelector_;
})



   

    describe("test if mint nft success",async function(){
        it("test nft owner is  minter",async function () { 
            let tx = await nft.mintNFT(deployer,"ipfs://bafkreiht5eqy3m3wjs4n67njkfysdytuh4e7q6ljjpiodz2nlbshdz2cz4");
            const owner = await nft.ownerOf(1);

            console.log(`totalSupply:${await nft.totalSupply()}`);
            expect(owner).to.be.equal(deployer);
            });
    });
    

    describe("test nft can be locked and send to dest chain",
        async function(){
            it("test send nft to dest chain,the nft should be locked on sourceNFT pool",async function(){
                    
                await ccipLocalSimulator.requestLinkFromFaucet(sourceNFTPool.target,ethers.parseEther("10"));

                    await nft.approve(sourceNFTPool.target,1);
                    console.log(`deployer:${deployer},destNFTPool:${destNFTPool.target},chainselector:${chainselector}`);
                    const tx = await sourceNFTPool.lockAndSendNFT(deployer,1,destNFTPool.target,chainselector);

                    const owner = await nft.ownerOf(1);
                    expect(owner).to.be.equal(sourceNFTPool.target);  
            })

            it("test wnft can be minted ",async function(){ 
                const owner = await wnft.ownerOf(1);
                expect(owner).to.be.equal(deployer);
            })
        
    });

   

    describe("test wnft can be burned and unlocked",
        async function(){ 
            it("test wnft can be burned",async function(){
                await ccipLocalSimulator.requestLinkFromFaucet(destNFTPool.target,ethers.parseEther("10"));

                await wnft.approve(destNFTPool.target,1);
                await destNFTPool.burnAndSendNFT(deployer,1,sourceNFTPool.target,chainselector);
                const totalSupply = await wnft.totalSupply();
                expect(0).to.be.equal(totalSupply);
            })

            it("test nft can be unlocked",async function(){
               const owner = await nft.ownerOf(1);
               expect(owner).to.be.equal(deployer);
            })
      
    }) 

