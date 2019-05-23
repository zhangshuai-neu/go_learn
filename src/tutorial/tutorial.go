package main

import ("fmt"
	"tutorial/testJson"
)

func main() { // { 不能在单独的行上
	fmt.Println("Hello, Go!")
/*
	fmt.Println("===== 1. 测试类型零值 =====")
	typeZeroVal.TestTypeZeroVal()

	fmt.Println("===== 2. 测试变量定义 =====")
	varUse.TestVarUse()

	fmt.Println("===== 3. 测试常量使用 =====")
	constUse.TestConstUse()

	fmt.Println("===== 4. 测试++,--和指针 =====")
	ptr.TestPtr()

	fmt.Println("===== 5. 测试条件表达式 =====")
	ifExpression.TestIfExpression()

	fmt.Println("===== 6. 测试循环表达式 =====")
	forExpression.TestForExpression()

	fmt.Println("===== 7. 测试switch表达式 =====")
	switchExpression.TestSwitchExpression()

	fmt.Println("===== 8. 跳转测试 =====")
	jumpLabel.TestJumpLabel()

	fmt.Println("===== 9. 数组测试 =====")
	array.TestArray()

	fmt.Println("===== 10. slice测试 =====")
	testSlice.TestSlice()

	fmt.Println("===== 11. channel测试 =====")
	testChannel.TestChannel()

	fmt.Println("===== 12. goroutine测试 =====")
	testGoroutine.TestGorourine()

	fmt.Println("===== 13. struct测试 =====")
	testStruct.TestStruct()

	fmt.Println("===== 14. 左值函数测试 =====")
	testFunc.TestFunc()

	fmt.Println("====== 15. 测试string =====")
	testString.TestString()

	fmt.Println("====== 16. 测试time =====")
	testTime.TestTime()
	*/

	fmt.Println("====== 17. 测试json=====")
	testJson.TestJson()
}