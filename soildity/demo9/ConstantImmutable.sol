// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract ConstantImmutable {
    string constant name = "Biden";  // 寫入棧中   只能使用 pure  (編譯后不可修改)
    uint immutable age;  // 初始化后不可修改  (部署合約后不可修改)

    constructor() {
        age = 100;
    }

    function getName() public pure returns(string memory) {
        return name;
    }

    function getAge() public view returns(uint) {
        return age;
    }
}