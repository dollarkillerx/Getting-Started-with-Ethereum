// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract ArrayExample {
    uint[] iArray;
    uint[] iArray2 = [1,2,3];
    uint[3] iArray3;

    function getArray() public view returns(uint []memory) {
        return iArray;
    }

    function getArrayByIndex(uint _i) public view returns(uint) {
        return iArray[_i];
    }

    function getLength() public view returns(uint) {
        return iArray.length;
    }

    function push(uint _i) public {
        iArray.push(_i);
    }

    // 刪除是變成0
    function delArray(uint _index) public {
        delete iArray[_index];
    }
}