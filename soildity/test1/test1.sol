pragma solidity 0.4.24;

contract TestMsgSender {
    address public owner;
    function get_owner() public {
        owner = msg.sender;
    }


    function get_money() public payable {

    }

    uint256 public bla;

    function test() public {
        bla = address(this).balance;
    }
}