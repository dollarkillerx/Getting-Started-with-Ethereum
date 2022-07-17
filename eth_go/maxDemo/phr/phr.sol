// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract TestA {

    string name;
    function getName() public view returns(string memory) {
        return name;
    }

    function setName(string memory _name) public {
        name = _name;
    }
}