// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract DollarKillerChain is ERC20, Ownable {
    bool public isLocked = false;

    // 在創建合約的1min 只能不能轉賬
    uint public timeLock = block.timestamp + 1 minutes;

    constructor() ERC20("DollarKillerChain","DKC") {
        _mint(msg.sender,100000 * 10 ** decimals());
    }

    function transfer(address _to, uint256 _amount) public override returns(bool) {
        // require(isLocked == false,"Transfer was locked");

        // require(block.timestamp > timeLock,"Is's not time yet");

        return super.transfer(_to,_amount);
    }

    function setLock() public onlyOwner returns(bool) {
        isLocked = !isLocked;

        return true;
    }

    // ERC 20 的有一個轉賬前置方法
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual override {
        require(isLocked == false,"Locked");
    }
}