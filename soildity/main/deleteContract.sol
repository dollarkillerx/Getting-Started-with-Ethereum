// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

// 向智能合约发送代币

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ReceiveTokenContract {
    IERC20 myToken;
    address owner;

    constructor(address _tokenAddress) {
        myToken = IERC20(_tokenAddress);
        owner = msg.sender;
    }

    // 注入流動性 
    function transferFrom(uint _amount) public {
        myToken.transferFrom(msg.sender,address(this),_amount); // 發送發,地址,金額
    }

    // 獲取餘額
    function getBalance(address _address) public view returns(uint) {
        return myToken.balanceOf(_address);
    }

    // 摧毀智能合約
    function kill() public {
        // 銷毀前把錢轉走
        myToken.transfer(owner,myToken.balanceOf(address(this)));
        // 銷毀合約
        selfdestruct(payable(owner));
    }

}