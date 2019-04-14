package varUse

import "fmt"

func TestVarUse(){
	var a,b,c int = 1,2,3
	var d = 4 	// 类型推断
	var s string = "VarUse"

	fmt.Println(a,b,c)
	fmt.Println(d)
	fmt.Println(s)
}