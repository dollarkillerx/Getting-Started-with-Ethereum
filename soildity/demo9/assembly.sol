// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

// https://solidity-cn.readthedocs.io/zh/develop/assembly.html

contract Assem {
    function add(uint a,uint b) public pure returns(uint) {
        assembly {
            let sum := add(a,b)
            mstore (0x0,sum) // 存儲到内存
            return (0x0,32)
        }
    }

    function call(uint a) public pure returns(uint) {
        bytes4 sig = bytes4(keccak256("call(uint)"));

        assembly {
            mstore(0x0,sig)
            let addressData := add(0x0,4)
            mstore(addressData,a)
            return (addressData,32)
        }
    }
}