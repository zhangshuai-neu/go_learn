package constUse

import "fmt"

const ci int = 10
const cs string = "cs"
const (
	RED int = iota //从0开始，常量计数器(记录常量的个数)
	BLACK
)
const (
	B  int64 = 1 << (iota * 10) // 1<<0
	KB                          // 1<<10
	MB                          // 1<<20
	GB                          // 1<<30
)

func TestConstUse() {

	fmt.Println("const int:", ci)
	fmt.Println("const string:", cs)

	fmt.Println("enum RED:", RED)
	fmt.Println("enum BLACK:", BLACK)

	fmt.Println("enum B:", B)
	fmt.Println("enum KB:", KB)
	fmt.Println("enum MB:", MB)
	fmt.Println("enum GB:", GB)
}
