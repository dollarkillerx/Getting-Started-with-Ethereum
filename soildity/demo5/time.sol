pragma solidity 0.4.24;

contract TestTime {

    function test1() public pure returns (string)  {
        if (1 minutes == 60 seconds) {
            return "1分鐘 = 60秒";
        }

        return "";
    }
}