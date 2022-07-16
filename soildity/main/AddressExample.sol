// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract AddressExample {
    address ownerAddress;

    constructor() {
        ownerAddress = msg.sender;
    }

    // 合約本身的地址
    function getContractAddress() external view returns (address) {
        return address(this);
    }

    // 調用智能合約人的地址
    function getSenderAddress() external view returns (address) {
        return address(msg.sender);
    }

    // 創建智能合約人的地址
    function getOwnerAddress() external view returns (address) {
        return ownerAddress;
    }

    function getBlance() external view returns(uint,uint,uint) {
        // 合約餘額度
        uint contractBalance = address(this).balance;
        // 調用賬戶餘額
        uint senderBalance = msg.sender.balance;
        // 創建者餘額
        uint ownerBalance = ownerAddress.balance;

        return (contractBalance,senderBalance,ownerBalance);
    }
}