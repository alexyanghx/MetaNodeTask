// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
import "@openzeppelin/contracts/access/Ownable.sol";
import "hardhat/console.sol";

contract BeggingContract is Ownable{


    mapping(address=>uint) private _donateRecords;

    event Donate(address indexed from, uint amount);

    function donate() public payable{
        _donateRecords[msg.sender] += msg.value;
        emit Donate(msg.sender, msg.value);
    }
    
    function withdraw() public onlyOwner{
        payable(msg.sender).transfer(address(this).balance);
    }

    function getDonation(address addr) public view returns(uint){
        return _donateRecords[addr];
    }
}