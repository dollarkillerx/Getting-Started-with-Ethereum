// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract Logic {
    uint private number;
    function setNumber() public {
        number += 1;
    }

    function getNumber() public view returns (uint) {
        return number;
    }
}