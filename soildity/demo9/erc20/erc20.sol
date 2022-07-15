// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

// 向智能合约发送代币

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract DollarkillerChain is ERC20 {
    constructor() ERC20("DollarKillerChain","DKL") {
        _mint(msg.sender,1000*10**decimals());
    }
}