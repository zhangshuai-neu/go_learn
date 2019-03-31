package main

import "fmt"

// my package
import (
    "tutorial/typeZeroVal"  //1. 测试类型零值
)

func main(){ 
    // { 不能在单独的行上
    fmt.Println("Hello, Go!")
    
    //1. 测试类型零值
    typeZeroVal.TestTypeZeroVal()
}
