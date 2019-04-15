package switchExpression

import "fmt"

func TestSwitchExpression(){
	var a int = 1
	switch a {
		case 1:
			fmt.Println("Case 1");
		case 2:
			fmt.Println("Case 2");
		default:
			fmt.Println("Default");
	}
}
