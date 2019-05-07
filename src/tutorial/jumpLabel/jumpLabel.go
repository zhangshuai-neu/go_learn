package jumpLabel

import "fmt"

func TestJumpLabel() {
	var i int
	var j int

	fmt.Println("jumpFor")
jumpBreak:
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if i == 1 {
				break jumpBreak // 跳出jumpFor指示的多级循环
			}
			fmt.Println("ij:", i*10+j)
		}
	}

	fmt.Println("jumpContinue")
jumpContinue:
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if i == 1 {
				continue jumpContinue // 结束这次的jumpContinue指示的多级循环
			}
			fmt.Println("ij:", i*10+j)
		}
	}

	fmt.Println("jumpGoto")
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if i == 1 {
				goto jumpGoto // goto的正常用法
			}
			fmt.Println("ij:", i*10+j)
		}
	}
jumpGoto:
}
