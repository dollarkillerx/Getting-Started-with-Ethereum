pragma solidity 0.4.24;

// inherit 繼承

contract Person {
    function say() public pure returns (string) {
        return "say hello";
    }
}

contract Job {
    function say() public pure returns (string) {
        return "job hello";
    }
}

// 繼承相同方法時   最遠繼承原則  (所以當前say方法是Job的)
contract Teacher is Person,Job {
    string public sy;
    function test() public {
        sy = say();
    }

}

contract Student  is Teacher {
    string public sy;
    function test() public {
        sy = say();
    } 
}