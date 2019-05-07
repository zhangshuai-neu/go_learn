package testFunc

import (
	"fmt"
)

type Int int

func (a *Int) add(b Int) {
	*a += b
}

func TestFunc(){
	var a *Int = new(Int)
	*a = 1
	a.add(1)
	fmt.Printf("%v\n",*a)
}