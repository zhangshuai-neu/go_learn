package testSlice

import "fmt"

func TestSlice() {
	// 1. 定义切片
	var s1 []int
	// 2. 向切片添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2, 3, 4)
	fmt.Println(s1)

	// 3. slice 定义和初始化
	s2 := []int{1, 2, 3}
	// 4. 直接打印
	fmt.Println(s2)
	// 5. 索引访问
	fmt.Println(s2[0])

	// 6. 使用make创建
	var s3 []int = make([]int, 3, 5)
	// 7. slice 的长度和容量
	s3[1] = 1
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s3), cap(s3), s3)

	// 7. slice为空
	var s4 []int
	if s4 == nil {
		println("s4 is nil")
	}

	// 8. 切片访问
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numbers[2:5]) // 左闭右开区间
	fmt.Println(numbers[:2])  // 左闭右开区间
	fmt.Println(numbers[2:])  // 左闭右开区间

	// 9. slice 拷贝
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	numbers1[0] = 100
	fmt.Println(numbers1)
	fmt.Println(numbers)

}
