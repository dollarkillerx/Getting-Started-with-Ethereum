// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

// Library類庫
/**

1. 不能被繼承
2. 不能包含鏈上數據
3. 不能接受代筆
4. 不能有構造函數
5. 部署一次可以多次使用

**/

library SafeMath {
    function add(uint a,uint b) public pure returns(uint) {
        uint c = a + b;
        require(c >= a);
        return c;
    }

    function sub(uint a,uint b) public pure returns(uint) {
        require(b <= a);

        return a - b;
    }
}

contract Example {
    using SafeMath for uint;

    function doAdd(uint _a,uint _b) public pure returns(uint) {
        return _a.add(_b);
    }

    function doSub(uint _a,uint _b) public pure returns(uint) {
        return _a.sub(_b);
    }
}