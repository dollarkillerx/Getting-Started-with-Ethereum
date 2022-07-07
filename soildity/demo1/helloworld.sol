pragma solidity 0.4.24;

contract Hello {
    string name; // 狀態變量 狀態變量默認為privet


    function test() public pure returns (int256)  {  // 函數默認是public
        int256 age;    // 局部變量 必須priver
        
        return age;
    }
}