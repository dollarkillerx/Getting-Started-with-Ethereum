// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;


// 銀行
contract Bank {

    mapping(address => uint256) bill; // 賬單 

    // 存錢 
    function saveMoney() public payable {
        bill[msg.sender] = bill[msg.sender] + msg.value;
    }

    // 查詢餘額
    function checkBalances() public view returns(uint256) {
        return bill[msg.sender];
    }

    // 獲取銀行準備金額
    function bankBalances() public view returns(uint256) {
        return address(this).balance;
    }

    // 取錢  (當前調用是合約 往 目標地址入金)
    function withdrawMoney(uint256 amount) public returns(bool) {
        require(bill[msg.sender] >= amount,"Insufficient balance");

        // (bool isSend,) = msg.sender.call{value: amount,gas: 5000}("");
        (bool isSend,) =payable(msg.sender).call{value: amount,gas: 5000}("");
        bill[msg.sender] =  bill[msg.sender] - amount;
        return isSend;
    }

    // 内部轉賬
    function internalTransfer(address _to, uint256 amount) public {
        require(bill[msg.sender] >= amount,"Insufficient balance");
        bill[msg.sender] =  bill[msg.sender] - amount;
        bill[_to] =  bill[_to] + amount;
    }

    // 如此調用就是 先 入金到 合約 然後再 通過合約賬戶 轉賬給 目標
    function publicTransfer(address payable _to) public payable returns(bool) {
        (bool isSend,) = _to.call{value: msg.value,gas: 5000}("");
        return isSend;
    }
}