// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./import1.sol";
// import "https://xxxxx.sol"  網絡上文件也可以

contract ImportExample2 {
    ImportExample importExample = new ImportExample();

    function getAge() public view returns (uint) {
        return importExample.age();
    }

    function getName() public view returns (string memory) {
        return importExample.getName();
    }
}