# Solidity

### Hello world

``` 
pragma solidity 0.4.24;                // 版本號

contract Hello {
    string name;                        // 狀態變量 狀態變量默認為privet


    function test() returns (int256) {  // 函數默認是public
        int256 age;                     // 局部變量 必須priver
        
        return age; 
    }
    
    func test2() {
        int256 age;
    }
    
    function test() public pure returns (int256)  {  // 函數默認是public
        int256 age;    // 局部變量 必須priver
        
        return age;
    }
}

//声明版本号
pragma solidity ^0.4.16;
//合约 有点类似于java中的class
contract HelloWorld{
    //合约属性变量
    string myName = "HelloWorld";
    //合约中方法 注意语法顺序 其中此处view 代表方法只读 不会消耗gas
    function getName() public view returns(string){
        return myName;
    }
    //可以修改属性变量的值 消耗gas
    function changeName(string _newName) public{
        myName = _newName;
    }
    // pure:不能读取也不能改变状态变量
    function pureName(string _name) public pure returns(string){
        return _name;
    }
    
}
```

#### 函数的访问权限：

- public 最大的访问权限，子类可以继承、可以访问，当前类能访问
- private 仅限内部访问，子类不能继承、不能访问
- internal 子类可以继承、可以访问，当前类可以访问
- external 子类可以继承、可以访问，当前类不能访问

#### 状态变量的访问权限：

- view/constant 对状态变量只读，这里的状态变量还包含区块链的内建对象数据、时间戳等  (只讀)
- pure 既不修改，也不读取状态变量的值 (不使用合約任何狀態變量)
- payable 調用函數需要付錢,錢付給了智能合約賬戶
- returns 返回值函數聲明
- external 僅合約外部可以調用，合約内部需要使用this調用
- internal 僅合約内部和繼承合約可以調用
- 訪問并且修改了變量 不使用修飾符

``` 
function test() public payable returns (uint256){
    // return this.balance;  // 獲取合約餘額
    return address(this).balance;
}
```

### 修飾符

變量：
- public 
- private  (默認)

函數：
- public (默認)
- private

#### 數據類型

``` 
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
```