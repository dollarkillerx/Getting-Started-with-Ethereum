pragma solidity 0.4.24;

contract TestModi {
   
    string name = "hellen";  // 默認 private
    
    function add() private { // 默認 public

    }

    // view 是用到的外部的變量
    function get_name() public view returns (string) {
        return name;
    }

    function test() public payable returns (uint256){
        // return this.balance;  // 獲取合約餘額
        return address(this).balance;
    }
}