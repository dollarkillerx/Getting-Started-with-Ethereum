// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**

Send:
Transfer:

上面不被推薦

Call:  屬於級別比較底的調用

**/

contract PayExample {

    // 用戶轉給 _to
    function send(address payable _to) public payable {
        bool isSend = _to.send(msg.value);
        require(isSend,"Send Fail");
    }

    function transfer(address payable _to) public payable {
        _to.transfer(msg.value);
    }

    function calls(address payable _to) public payable {
        // (bool isSend,) = _to.call{value: msg.value}("");
        (bool isSend,) = _to.call{value: msg.value,gas: 5000}("");
        require(isSend,"Send Fail");
    }
}