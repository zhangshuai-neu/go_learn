package testString

import (
	"fmt"
	"strconv"
)

func TestString(){
	// string 比较大小
	s1 := "123"
	s2 := "21"
	s3 := "1234"
	s1_n, _:= strconv.ParseInt(s1, 10, 64)
	s2_n, _:= strconv.ParseInt(s2, 10, 64)
	s3_n, _:= strconv.ParseInt(s3, 10, 64)
	fmt.Println(s1,s2, s1_n < s2_n )
	fmt.Println(s2,s3, s2_n < s3_n )
}