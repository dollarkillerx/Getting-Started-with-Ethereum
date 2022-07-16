// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract EnumExample {
    enum Status {OFF,ON}

    Status status;

    function getStatus() view external returns (Status) {
        return status;
    }

    function setStatus(Status _status) external {
        status = _status;
    }

    function setStatusNo() external {
        status = Status.ON;
    }

    function reset() external {
        delete status; // 刪除枚舉
    }

    function getKeyByValue(Status _status) pure external returns (string memory) {
        if (_status == Status.ON) {
            return "No";
        }

        return "OFF";
    }

    function getValueByKey(string memory _status) pure external returns (Status) {
        if (keccak256(bytes(_status))  == keccak256(bytes("OFF"))) return Status.OFF;

        return Status.ON;
    }
}