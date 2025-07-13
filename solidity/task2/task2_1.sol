// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// 合约包含以下标准 ERC20 功能：
// balanceOf：查询账户余额。
// transfer：转账。
// approve 和 transferFrom：授权和代扣转账。
// 使用 event 记录转账和授权操作。
// 提供 mint 函数，允许合约所有者增发代币。
// 提示：
// 使用 mapping 存储账户余额和授权信息。
// 使用 event 定义 Transfer 和 Approval 事件。
// 部署到sepolia 测试网，导入到自己的钱包
contract TestToken is Ownable{
    using SafeMath for uint256;

    mapping(address => uint) private _balances;

    mapping(address holder=>mapping(address spender=>uint amount)) private _allowances;

    uint private _totalSupply;

    event Transfer(address from ,address to,uint amount);
    event Approval(address owner , address spender , uint amount);

    constructor() Ownable(msg.sender){}

    function balanceOf(address addr) public view returns(uint){
        return _balances[addr];
    }

    function allowance(address owner,address speeder) public view returns(uint){
        return _allowances[owner][speeder];
    }

    function _approve(address owner,address spender,uint amount) internal {
        require(owner!=address(0),"approve owner the zero address");   
        require(spender!=address(0),"approve spender the zero address");    
        _allowances[msg.sender][spender]=_allowances[msg.sender][spender].add(amount);
        emit Approval(msg.sender, spender, amount);
    }

    function approve(address spender,uint amount) public {
        _approve(msg.sender, spender, amount);
    }

    function transerFrom(address sender,address reciept,uint amount) public {
        require(sender!=address(0),"transerFrom sender the zero address");
        require(reciept!=address(0),"transerFrom reciept the zero address");

        _allowances[sender][msg.sender]=_allowances[sender][msg.sender].sub(amount,"not enough allowance");
        _transfer(sender, reciept, amount);
    }


    function _transfer(address from,address to,uint amount) internal {
        require(from!=address(0),"transfer from the zero address");
        require(to!=address(0),"transfer to the zero address");

        _balances[from]=_balances[from].sub(amount,"balance not enough");
        _balances[to]=_balances[to].add(amount);
        emit Transfer(from, to, amount);
    }

    function transfer(address reciept,uint amount) public {
        _transfer(msg.sender, reciept, amount);
    }


    function mint(address account,uint amount) public onlyOwner {
        require(account!=address(0),"mint the zero address");

        _totalSupply = _totalSupply.add(amount);
        _balances[account] = _balances[account].add(amount);
        emit Transfer(address(0),account, amount);

    }

}
