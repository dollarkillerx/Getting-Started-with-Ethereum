# Solidity

### Tools
- Ganache
- Remix

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
- external 子类可以继承、可以访问，当前类不能访问    (只能外部調用)

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

#### 地址  (地址是合約的基礎)

- balance: 獲取餘額
- transfer: 轉賬
- send: 轉賬不安全 不推薦使用， 合約餘額不夠需要自己手動處理 不會報錯 [慎用]
- call: 合約内部調用合約，調用底層代碼 [慎用]
- callcode: 調用底層代碼 [慎用]
- delegatecall: 調用底層代碼 [慎用]

``` 
    uint256 public thisB;
    function tesx() {
        thisB = address(this).balance;  // this 指向合約本身
    }

    // payable 賬戶轉賬到合約
    function constract_get_money() public payable {

    }

    // 獲取合約餘額
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
```


#### 高級語法部分:

- `msg.sender` `address`  調用當前合約的 的地址
- `msg.value`  `uint(單位wei)` 是調用當前 `payable` 方法傳入的金額
- `msg.data` `bytes` 完整調用數據

- `block.number` `uint` 當前區塊號
- `blockhash` `byte32` 當前區塊hash
- `block.gaslimit` `uint` 當前的gaslimit
- `block.timestamp` `uint` 當前區塊的時間戳

- `now` 當前區塊的時間戳
- `gasleft()` `uint` 剩餘的gas
- `tx.gasprice` `uint` 交易 `gas` 價格

#### 錯誤處理方式:

- `require(msg.value == 10)`  如果不滿足條件就會報錯   (不會消耗gas)  (官方推薦)
- `assert(msg.value == 10)`   如果不滿足條件就會報錯   (會消耗gas)
- `revert()`  處理複雜的邏輯  更加靈活多變

``` 
if (msg.value != 10) {
    revert();
}

money = msg.value
```

- `throw` 以廢棄

``` 
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
```


#### 時間

- 1 minutes == 60 seconds
- 1 hours == 60 minutes
- 1 day == 24 hours
- 1 weeks == 7 days
- 1 years == 365 days

``` 
   function test1() public pure returns (string)  {
        if (1 minutes == 60 seconds) {
            return "1分鐘 = 60秒";
        }

        return "";
    } 
```

##### 訪問函數 & 調用其他合約

``` 
    string name = "hello";
    
    string public new_name = "hello2";

    function get_name() public view returns (string) {
        return name;
    }


}

contract TestFunc2 {
    function getTestFuncNewName() public returns (string) {
        TestFunc t = new TestFunc();  // 聲明并創建
        return t.new_name();
    }
    function getTestFuncNewName2() public returns (string) {
        TestFunc public tf;   // 聲明
        address addr = new TestFunc(); // 獲取地址
        tf = TestFunc(addr);  // 賦值
        return tf.new_name();
    }
}
```

##### 合約之間的轉賬

``` 

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
```

##### 繼承

- is 繼承
- 繼承相同方法時   最遠繼承原則  (所以當前say方法是Job的)

``` 
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
contract Teacher is Person,Job {  // is 繼承
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
```

- Modifier 修飾器

``` 
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
```