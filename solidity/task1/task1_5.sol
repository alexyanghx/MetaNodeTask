// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
import "hardhat/console.sol";

contract MergArr{

    function mergArr(int[] calldata nums1,int[] calldata nums2) public pure returns(int[] memory){
        uint i=0;
        uint j=0;
        uint k=0;

        int[] memory nArr = new int[](nums1.length+nums2.length);

        while(i<nums1.length && j<nums2.length){
            if(nums1[i]<nums2[j]){
                console.log(1);
                nArr[k] = nums1[i];
                i+=1;
            }else{
                console.log(2);
                nArr[k] = nums2[j];
                j+=1;
            }
            k+=1;
        }

        
        while(i<nums1.length){
            nArr[k] = nums1[i];
            i+=1;
            k+=1;
        }
        


        while(j<nums2.length){
            nArr[k] = nums2[j];
            j+=1;
            k+=1;
        }
        

        return nArr;

    }
}
