package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	First string
	Last string
	Age int
}
func main(){
	j:=`[{"First":"james","Last":"bond","Age":32},{"First":"ansh","Last":"me","Age":34}]`
	s:=[]byte(j)
	p:= []people{}
	err:=json.Unmarshal(s,&p)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	}
