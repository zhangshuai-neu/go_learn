package main

import "fmt"

// my package
import (
    "tutorial/typeZeroVal"  //1. 测试类型零值
    "tutorial/varUse"
    "tutorial/constUse"
    "tutorial/ptr"
)

func main(){  // { 不能在单独的行上
    fmt.Println("Hello, Go!")
    
    fmt.Println("===== 1. 测试类型零值 =====")
    typeZeroVal.TestTypeZeroVal()


    fmt.Println("===== 2. 测试变量定义 =====")
    varUse.TestVarUse()

    fmt.Println("===== 3. 测试常量使用 =====")
    constUse.TestConstUse()

    fmt.Println("===== 4. 测试++,--和指针 =====")
    ptr.TestPtr()

}
