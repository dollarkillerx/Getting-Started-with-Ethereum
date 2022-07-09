pragma solidity 0.4.24;

// modifier 修飾器

contract TestModifier {

    address admin;
    // 構造函數
    constructor() public {
        admin = msg.sender;
    }

    // 修飾器
    modifier auth_filter() {
        // 校驗是不是管理員
        require(msg.sender == admin);
        _;
    }

    // 權限校驗,只能部署者 也就是管理員調用,其他用戶不行
    function get_name() public view auth_filter returns (string) {
        return "hallen";
    }

}