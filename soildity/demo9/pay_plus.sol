// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract Reciver {
    event Recived(address caller, uint amount, string message);

    // receive 儅賺錢進來時 就會觸發
    receive() external payable {
        emit Recived(msg.sender,msg.value,"Receive was called");
    }

    // 儅轉裝 含有數據時 _addr.call{value: msg.value}("數據");    下面fallback 就會被調用
    fallback() external payable {
         emit Recived(msg.sender,msg.value,"Fallback was called");
    }

    function foo(string memory _message,uint _y) public payable returns (uint) {
        emit Recived(msg.sender,msg.value,_message);
        return _y;
    }

    function getAddress() public view returns(address) {
        return address(this);
    }

    function getBalance() public view returns(uint) {
        return address(this).balance;
    }
}

contract Caller {
    Reciver receiver;

    constructor() {
        receiver = new Reciver();
    }

    event Response(bool success, bytes data);

    // 調用receive方法 因爲沒有msg.data
    function testCall(address payable _addr, uint _y) public payable {
        // (bool success,bytes memory data) = _addr.call{value: msg.value}("");
        (bool success,bytes memory data) = _addr.call{value: msg.value}(
            abi.encodeWithSignature("foo(string,uint256)","call foo",_y)
        );
        emit Response(success,data);
    }

    function getAddress() public view returns (address) {
        return receiver.getAddress();
    }

    function getBalance() public view returns(uint) {
        return receiver.getBalance();
    }
}