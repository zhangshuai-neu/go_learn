package forExpression

import "fmt"

func TestForExpression(){
	// 1. 无限循环
	fmt.Println("1. 无限循环");
	var i int = 1;
	for {
		if i>=3 {
			break
		}
		fmt.Println("i:",i);
		i++
	}
	

	// 2. 带有条件表达式
	fmt.Println("2. 带有条件表达式");
	i = 1
	for i<3 {
		fmt.Println("i:",i);
		i++
	}

	// 3. 经典形式
	fmt.Println("3. 经典形式");
	for i=1; i<3; i++ {
		fmt.Println("i:",i);
	}
}