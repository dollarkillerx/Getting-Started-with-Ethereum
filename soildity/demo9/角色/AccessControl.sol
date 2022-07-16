// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract MyToken is ERC20, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    constructor() ERC20("MyToken", "MTK") {
        _mint(msg.sender, 100000000 * 10 ** decimals());
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);

        // 創建一個 MINTER_ROLE 的權限組
        _grantRole(MINTER_ROLE, msg.sender);
    }

    function mint(address to, uint256 amount) public onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }
}

// MINTER_ROLE 擁有權限者
// onlyRole   修飾器 只能 onlyRole 調用
// hasRole 查看用戶是否有role  (roleID,userID)
// grantRole 分配role
// revokeRole 銷毀role
// renounceRole 放棄role