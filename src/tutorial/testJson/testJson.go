package testJson

import (
	"encoding/json"
	"fmt"
)
type Info struct {
	Switch    string `json:"switch"`
	WhiteList map[string]bool `json:"whiteList"`
}

func TestJson(){
	jsonStr := "{\"switch\": \"1\",\"whiteList\": {\"92586523228\":true,\"123123\":true} }"
	info := Info{"", map[string]bool{}}
	json.Unmarshal( []byte(jsonStr), &info )
	fmt.Println(info)
	fmt.Println(info.WhiteList["92586523228"])
	fmt.Println(info.WhiteList["123123"])
}
