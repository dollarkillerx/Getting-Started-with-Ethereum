pragma solidity 0.4.24;

contract TestModi {
   
    string name = "hellen";  // 默認 private
    
    function add() private { // 默認 public

    }

    // view 是用到的外部的變量
    function get_name() public view returns (string) {
        return name;
    }

    function getMoney() public payable returns (uint256){
        // return this.balance;  // 獲取合約餘額
        return address(this).balance;
    }


    uint256 public thisB;
    function tesx() {
        thisB = address(this).balance;
    }

    // payable 賬戶轉賬到合約
    function constract_get_money() public payable {

    }

    // 獲取合約金額
    uint256 public my_balance;
    function get_constract_balance() public {
       my_balance =  address(this).balance;
    }

    // 轉賬目標地址
    address addr = 0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2;
    // 轉賬到用戶
    function trans() public {
        addr.transfer(5 * 10 ** 18); // 單位為 wei
    }

}