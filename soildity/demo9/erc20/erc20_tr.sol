// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

// 向智能合约发送代币

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ReceiveTokenContract {
    IERC20 myToken;

    constructor(address _tokenAddress) {
        myToken = IERC20(_tokenAddress);
    }

    // 注入流動性 
    function transferFrom(uint _amount) public {
        myToken.transferFrom(msg.sender,address(this),_amount); // 發送發,地址,金額
    }

    // 獲取餘額
    function getBalance(address _address) public view returns(uint) {
        return myToken.balanceOf(_address);
    }
}

// 注意轉賬給合約 需要 先調用 ERC20 的 approve 授權給該合約 才能轉賬