// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract A {
    function callBFunction(address _address,uint256 _num,string memory _message) public returns (bool) {
        (bool success,) = _address.call(
            abi.encodeWithSignature("bFunction(uint256,string)",_num,_message)
        );

        return success;
    }

    function callBFunction2(address _address,uint256 _num,string memory _message) public returns (bool) {
        bytes4 sig = bytes4(keccak256("bFunction(uint256,string)"));
        bytes memory _bNum = abi.encode(_num);
        bytes memory _bMessage = abi.encode(_message);

        (bool success,) = _address.call(
            abi.encodePacked(sig,_bNum,_bMessage)
        );
        return success;
    }


    function callBFunction3(address _address,uint256 _num,string memory _message) public returns (bool) {
        bytes4 sig = bytes4(keccak256("bFunction(uint256,string)"));
        (bool success,) = _address.call(
            abi.encodeWithSelector(sig,_num,_message)
        );
        return success;
    }

    function callBFunction4(address _address,uint256 _num,string memory _message) public returns (bool) {
        B cb = B(_address);
        cb.bFunction(_num,_message);
        return true;
    }
}

contract B {
    uint256 public num;
    string public message;

    function bFunction(uint256 _num, string memory _messsge) public returns (uint256, string memory) {
        num = _num;
        message = _messsge;
        return (num,message);
    }
}

