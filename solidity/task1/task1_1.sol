// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


//✅ 创建一个名为Voting的合约，包含以下功能：
// 一个mapping来存储候选人的得票数
// 一个vote函数，允许用户投票给某个候选人
// 一个getVotes函数，返回某个候选人的得票数
// 一个resetVotes函数，重置所有候选人的得票数
contract Voting {

    mapping(address candidate  =>uint voteCount) votes;
    mapping(address voter=>bool voted) voteRecord;
    address[] candidates;


    function candidateExist(address addr) private view returns(bool){
        for(uint i=0;i<candidates.length;i++){
            if(addr==candidates[i]) return true;
        }
        return false;
    }

    function vote(address candidate) public{

        require(!voteRecord[msg.sender],"had voted!");

        voteRecord[msg.sender] = true;

        if(!candidateExist(candidate)){
            candidates.push(candidate);
        }

        votes[candidate]+=1;
    }

    function getVotes(address candidate) public view returns(uint){
        return votes[candidate];
    }

    function resetVotes()public {
        for(uint i=0;i<candidates.length;i++){
            votes[candidates[i]] = 0;
        }
    }
    
}
