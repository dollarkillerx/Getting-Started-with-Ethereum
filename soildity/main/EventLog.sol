// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract ExampleEvent {
    event Deposit(address _from, string _name, uint256 _value);

    function deposit(string memory _name) public payable {
        emit Deposit(msg.sender,_name,msg.value);
    }
}

/**

引用类型（指针传递）， 没有*号操作符，而是使用两个关键字来表示 `deposit(string memory _name)`

memory(值类型)
storage（引用类型）


日志會寫在區塊鏈中
**/