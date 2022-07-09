pragma solidity 0.4.24;

contract TestFunc {
    string name = "hello";
    
    string public new_name = "hello2";

    function get_name() public view returns (string) {
        return name;
    }


}

contract TestFunc2 {
    // TestFunc t = new TestFunc();
    
    function getTestFuncNewName() public returns (string) {
        TestFunc t = new TestFunc();
        return t.new_name();
    }
}

contract Trans1 {
    // 查看當前合約的餘額
    function get_balance() public view returns (uint256) {
        return address(this).balance;
    }
     // 想當前合約存入金額
    function contract_get_money() public payable {

    }
}

contract Trans2 {

    uint256 public tran1Addr;

    function tran1_addr() public {
        address addr = new Trans1(); // 810453375409437434270020959927630513759751760877   這樣會新實例化一個合約
        
        tran1Addr = uint256(addr); // 0x3328358128832A260C76A4141e19E2A943CD4B6D
    }

    // 查看當前合約的餘額
    function get_balance() public view returns (uint256) {
        return address(this).balance;
    }

    // 想當前合約存入金額
    function contract_get_money() public payable {

    }

    // 初始化 trans1
    Trans1 trans1;
    function getTrans1Addr(address addr) public {
        trans1 = Trans1(addr);
    }

    // 轉賬給Trans1
    function pay_to_t1() public {
        // 調用trans1轉賬方法  5eth  gas最高消耗800
        trans1.contract_get_money.value(5 * 10 ** 18).gas(800)();
    }

}