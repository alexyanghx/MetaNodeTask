// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 用 solidity 实现整数转罗马数字
contract Int2Roma{

    function int2roma(uint num) public pure returns(string memory ){
        string memory roman_num = "";
        string[13] memory romans = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"] ;
	    uint[13] memory nums = [uint(1000), 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];

        for(uint i=0;i<romans.length && num>0;){
            if(num>= nums[i]){
                roman_num = string.concat(roman_num, romans[i]);// 拼接字符
                num-=nums[i];
                continue;
            }
            i+=1;
        }

        return roman_num;
	    
    }

    

}
