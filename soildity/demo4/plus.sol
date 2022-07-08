pragma solidity 0.4.22;

// 高級語法部分

contract Plus {
    address public owner;
    // msg.sender 調用當前合約的 的地址
    function get_owner() public {
        if (owner== 0) {  // 這樣可以做到 owner之記錄首次調用者的地址
            owner = msg.sender;
        }
    }

    // 限定合規的錢才能轉賬
        // msg.value 是調用當前方法傳入的金額
    function contract_get_money() public payable returns(string) {

        // require(msg.value > 5 * 10 ** 18);
        if (msg.value < 5 * 10 ** 18) {
            revert();
            return "轉賬金額必須 >= 5 eth";
        }


        return "轉賬成功";
    }

    uint256 public contract_balance;
    function get_balance() public {
        contract_balance = address(this).balance;
    }
}