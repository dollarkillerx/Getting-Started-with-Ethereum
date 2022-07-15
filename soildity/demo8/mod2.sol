// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract ExampleModifier {
    address public owner; // 只能合約創建者
    uint256 public account;

    constructor() {
        // 初始化變量
        owner = msg.sender;
        account = 0;
    }

    modifier authFilter() {
        require(msg.sender == owner,"Only Owner"); // (條件/異常名稱)
        _;
    }

    modifier valid100(uint256 _account) {
         require(_account > 100,"valid100");
        _;
    }

    function updateAccount(uint256 _account) public {
        if (msg.sender == owner) {
            account = _account;
        }
    }

      function updateAccount2(uint256 _account) public authFilter valid100(_account) {
        if (msg.sender == owner) {
            account = _account;
        }
    }
}