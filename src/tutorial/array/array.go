package array

import "fmt"

func TestArray() {
	// 数组的初始化
	var arrayInt [2]int = [2]int{1, 2}
	var arrayInt1 [2]int = [2]int{1, 2}
	var arrayIntPtr *[2]int = &arrayInt

	fmt.Println("arrayInt[0]: ", arrayInt[0])
	fmt.Println("arrayInt[0]: ", arrayInt[1])
	fmt.Println("eqaul:", arrayInt == arrayInt1)

	(*arrayIntPtr)[0] = 3
	(*arrayIntPtr)[1] = 4
	fmt.Println("arrayInt[0]: ", arrayInt[0])
	fmt.Println("arrayInt[0]: ", arrayInt[1])
	fmt.Println("eqaul:", arrayInt == arrayInt1)

}
