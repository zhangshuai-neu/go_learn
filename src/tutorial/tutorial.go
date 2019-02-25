package main

import "fmt"

func test_var(){
    fmt.Println("\ntest_var ===========")
    var s string = "hello"
    var i int = 10
    fmt.Println(s,i)
}

func main(){ 
    // { 不能在单独的行上
    fmt.Println("Hello, Go!")
    test_var()

}
