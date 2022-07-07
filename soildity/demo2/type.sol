pragma solidity 0.4.24;

contract DataType {

    bool public ok = true;
    bool public fu;

    function test() public view returns (string) {
        if (fu) {
            return "is ok";
        }

        return "is not ok";
    }

    string public px = test();

   // # int

    int8 a = 1;
    int16 b = 3;

    function add() public view returns (int8) {
        return a + int8(b);
    }

    int8 public c = add();

    // # bytes

    bytes1 public b1 = "h";
    bytes2 public b2 = "he";
    bytes11 public b3 = "cxsad";
    // bytes長度不可以修改
    // 元素值不可修改,只讀
    
    function getLen() public view returns(uint256) {
        return b3.length;  // 返回時類型的長度，不是值的長度
    }

    function getByIndex()  public view returns(bytes1) {
        return b3[0];
    }

    // # 枚舉 Enums

    enum Gender {  // 枚舉為自定義類型
        male,
        female
    }

    Gender public defaultGender = Gender.male;

    // # 函數
    function testFunc(string name) public pure returns(string) {
        return name;
    }

    string public pxT1 = testFunc("hello world");


    // 多個返回值
    function pt2More(string name) public pure returns (string,string) {
        return (name,"hello");
    }


    // string public tx1;
    // string public tx2;
    
    // (tx1,tx2) = pt2More("dollarkiller");
    // (,tx2) = pt2More("dollarkiller");


     // 構造函數
    constructor() {
        
    }
}