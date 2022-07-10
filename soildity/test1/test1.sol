pragma solidity 0.4.24;

contract TestMsgSender {
    address public owner;
    function get_owner() public {
        owner = msg.sender;
    }
}