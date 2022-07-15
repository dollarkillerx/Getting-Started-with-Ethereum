// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;


// 定義error
error Unauthorized(string error);
// error Unauthorized(string error, address _address); 多個參數
// 調用 Unauthorized({error: "xxx", address: "xxx"});

contract ErrorExample {
    address payable owner = payable(msg.sender);

    function withraw() public {
        if(msg.sender != owner) {
            revert Unauthorized("Unatuorized");
        }

        owner.transfer(address(this).balance);
    }
}