package ptr

import "fmt"

func TestPtr(){
	var i int = 1
	var ptr *int = &i
	i--
	i++
	i++
	fmt.Println("i:",i)
	fmt.Println("ptr(point to i):",*ptr)
}
