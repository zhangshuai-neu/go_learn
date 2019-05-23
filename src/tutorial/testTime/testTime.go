package testTime

import (
	"fmt"
	"time"
)
func TestTime(){
	const baseFormat = "2006-01-02 15:04:05"

	var testTimeStr string = "2018-01-02 15:04:05"
	tt,_:= time.Parse(baseFormat,testTimeStr)
	fmt.Println(tt.Format(baseFormat))


	// 起始时间戳
	var startTimeStamp int64 = 1555171200
	st := time.Unix(startTimeStamp,0)
	fmt.Println(st.Format(baseFormat))

	fmt.Println(tt.Before(st))
}