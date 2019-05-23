package testTypeConversion

import (
	"fmt"
)

func TestTypeConversion() {
	var i int32
	var ui uint32
	i = -1
	ui = uint32(i)
	fmt.Printf("int:%v ,uint %v", i, ui)
}