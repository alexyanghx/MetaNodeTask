// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ReverseString{

    function reverse(string memory _str) public pure returns (string memory){
        bytes memory str_bytes = bytes(_str);
        uint length = str_bytes.length;
        bytes memory reversedStrBytes =new bytes(length);
        uint j =0;
        while(j<length){
            //中文
            if(str_bytes[j]&0x80!=0){
                reversedStrBytes[length-j-3] =str_bytes[j];
                reversedStrBytes[length-j-2] =str_bytes[j+1];
                reversedStrBytes[length-j-1] =str_bytes[j+2];
                j+=3;
            }else{
                reversedStrBytes[length-j-1] =str_bytes[j];
                j++;
            }

        }
        return string(reversedStrBytes);
    }
}
