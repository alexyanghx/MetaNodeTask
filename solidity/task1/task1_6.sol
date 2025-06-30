                                                                               // SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BinarySearch{

    function binarySearch(int[] calldata nums,int num) public pure returns(int index){

        uint left=0;
        uint right=nums.length-1;
        
        while(left<=right){
            uint mid = left +(right-left)/2;
            if(nums[mid]==num){
                return int(mid);
            }else if(nums[mid]<num){
                left = mid+1;
            }else{
                right=mid-1;
            }
        }

        return -1;
    }
}
