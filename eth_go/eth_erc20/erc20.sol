// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";  // 導入 openzeppelin


contract DollarKillerChain is ERC20 { // 繼承ERC20
    constructor() ERC20("DollarkillerChain","TKC") { // 初始化ERC20 (Token名稱,代幣名稱)
        _mint(msg.sender,100000000 * 10 ** decimals()); // 預 挖 1億個 代幣
    }
}
